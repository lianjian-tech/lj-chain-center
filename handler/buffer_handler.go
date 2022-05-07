package handler

import (
	"fmt"
	"lj-chain-center/common/util"
	"sync"
	"time"
)

var queryBuf sync.Map

type ifData struct {
	ts   time.Time
	Data interface{}
}

func GetQueryKey(ifName string, req interface{}) string {
	val1, ok := req.(int64)
	if ok {
		return fmt.Sprintf("%v:%v", ifName, val1)
	}
	val2, ok := req.(uint64)
	if ok {
		return fmt.Sprintf("%v:%v", ifName, val2)
	}
	val3, ok := req.(string)
	if ok {
		return fmt.Sprintf("%v:%v", ifName, val3)
	}
	return fmt.Sprintf("%v:%v", ifName, util.ToJSONStr(req))
}

func LoadBuffer(key string, ts int) interface{} {
	ret, ok := queryBuf.Load(key)
	if ok {
		data, ok := ret.(*ifData)
		if ok && data != nil && time.Since(data.ts) <= time.Duration(ts)*time.Second {
			//log.Infof("LoadBuffer data exist, key:%v", key)
			return data.Data
		}
	}
	//log.Infof("LoadBuffer data not exist, key:%v", key)
	return nil
}

func StoreBuffer(key string, data interface{}) {
	queryBuf.Store(key, &ifData{time.Now(), data})
}
