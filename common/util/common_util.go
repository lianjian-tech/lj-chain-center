package util

import (
	"bytes"
	"crypto/md5"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/url"
	"reflect"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

func BytesToInt(b []byte) (acc uint64) {
	length := len(b)
	if length%4 != 0 {
		extra := (4 - length%4)
		b = append([]byte(strings.Repeat("\000", extra)), b...)
		length += extra
	}
	var block uint32
	for i := 0; i < length; i += 4 {
		block = binary.BigEndian.Uint32(b[i : i+4])
		acc = (acc << 32) + uint64(block)
	}
	return
}

func ConvertStringToBoolean(booleanString string, defaultValue bool) bool {
	retBool := defaultValue

	if len(booleanString) > 0 {
		b, err := strconv.ParseBool(booleanString)
		if err == nil {
			retBool = b
		}
	}
	return retBool
}

func ConvertStringToFloat(val string, defaultValue float64) float64 {
	ret := defaultValue

	if len(val) > 0 {
		b, err := strconv.ParseFloat(val, 64)
		if err == nil {
			ret = b
		}
	}
	return ret
}

func ConvertStringToInt(intString string, defaultValue int) int {
	retInt := defaultValue

	if len(intString) > 0 {
		i64, err := strconv.Atoi(intString)
		if err == nil {
			retInt = i64
		}
	}
	return retInt
}

func ConvertStringToInt64(intString string, defaultValue int64) int64 {
	retInt := defaultValue

	if len(intString) > 0 {
		i64, err := strconv.ParseInt(intString, 10, 64)
		if err == nil {
			retInt = i64
		}
	}
	return retInt
}

func GetRandomString(n int) string {
	const symbols = "0123456789abcdefghjkmnopqrstuvwxyzABCDEFGHJKMNOPQRSTUVWXYZ"
	const symbolsIdxBits = 6
	const symbolsIdxMask = 1<<symbolsIdxBits - 1
	const symbolsIdxMax = 63 / symbolsIdxBits

	prng := rand.New(rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, n)
	for i, cache, remain := n-1, prng.Int63(), symbolsIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = prng.Int63(), symbolsIdxMax
		}
		if idx := int(cache & symbolsIdxMask); idx < len(symbols) {
			b[i] = symbols[idx]
			i--
		}
		cache >>= symbolsIdxBits
		remain--
	}
	return string(b)
}

func MD5(text string) string {
	ctx := md5.New()
	ctx.Write([]byte(text))
	return hex.EncodeToString(ctx.Sum(nil))
}

func IsMailFormat(mail string) bool {
	var mailRe = regexp.MustCompile(`\A[\w+\-.]+@[a-z\d\-]+(\.[a-z]+)*\.[a-z]+\z`)
	return mailRe.MatchString(mail)
}

func NowStr() string {
	tm := time.Unix(time.Now().Unix(), 0)
	return tm.Format("2006-01-02 15:04:05")
}

func GenTimeStr(ts int64) string {
	return time.Unix(ts, 0).Format("2006-01-02")
}

func IsEmptyStr(needCheck string) bool {
	if needCheck != "" && len(needCheck) > 0 {
		return false
	}
	return true
}

func ToJSONStr(val interface{}) string {
	if nil == val {
		return ""
	}
	real := reflect.ValueOf(val)
	if real.IsNil() {
		return ""
	}
	if real.Kind() == reflect.Ptr && !real.Elem().IsValid() {
		return ""
	}
	if (real.Kind() == reflect.Slice || real.Kind() == reflect.Array || real.Kind() == reflect.Map) && real.IsNil() {
		fmt.Printf("list:%#v\n", real)
		return ""
	}
	data, err := json.Marshal(val)
	if nil != err {
		return fmt.Sprintf("%#v", val)
	}
	return string(data)
}

func CatchError() {
	if err := recover(); err != nil {
		fmt.Errorf("recover error:%v\n", err)
	}
}

func GetDatetime(layout string) int64 {
	timeStr := time.Now().Format(layout)
	time, _ := time.Parse(layout, timeStr)
	datetime := time.Unix() - 60*60*8
	return datetime * 1000
}

//签名
func SignMd5(content interface{}, privateKey string) (string, error) {
	reqMap := ToMap(content)
	value := JoinStringsInASCII(reqMap, "sign")
	sign, err := RsaSignWithMd5(value, privateKey)
	if err != nil {
		return "", err
	}
	sign = url.QueryEscape(sign)
	return sign, nil
}

//验签
func VerifySignMd5(content interface{}, publicKey string) bool {
	reqMap := ToMap(content)
	fmt.Printf("reqMap 111:%v\n", ToJSONStr(reqMap))
	sign := reqMap["sign"].(string)
	fmt.Printf("sign:%v\n", sign)
	delete(reqMap, "sign")
	fmt.Printf("reqMap 222:%v\n", ToJSONStr(reqMap))
	value := JoinStringsInASCII(reqMap, "sign")
	sign, err := url.QueryUnescape(sign)
	if err != nil {
		return false
	}
	return RsaVerifySignWithMd5(value, sign, publicKey)
}

//参数名ASCII码从小到大排序后拼接 不包含空值
func JoinStringsInASCII(data map[string]interface{}, exceptKeys ...string) string {
	var list []string
	m := make(map[string]int)
	if len(exceptKeys) > 0 {
		for _, except := range exceptKeys {
			m[except] = 1
		}
	}
	for k := range data {
		if _, ok := m[k]; ok {
			continue
		}
		value := data[k]
		if value == "" {
			continue
		}
		list = append(list, fmt.Sprintf("%s=%s", k, value))
	}
	sort.Strings(list)
	return strings.Join(list, "&")
}

func ToMap(content interface{}) map[string]interface{} {
	var name map[string]interface{}
	if marshalContent, err := json.Marshal(content); err != nil {
		fmt.Println(err)
	} else {
		d := json.NewDecoder(bytes.NewReader(marshalContent))
		d.UseNumber() // 设置将float64转为一个number
		if err := d.Decode(&name); err != nil {
			fmt.Println(err)
		} else {
			for k, v := range name {
				name[k] = v
			}
		}
	}
	return name
}
