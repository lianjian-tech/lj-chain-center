package util

import (
	"crypto"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"strings"
)

//签名
func RsaSignWithMd5(data string, privateKey string) (string, error) {
	privateKey = Base64URLDecode(privateKey)
	key, err := base64.StdEncoding.DecodeString(privateKey)
	if err != nil {
		return "", err
	}
	prvKey, err := x509.ParsePKCS8PrivateKey(key)
	if err != nil {
		return "", err
	}

	hashMd5 := md5.Sum([]byte(data))
	hashed := hashMd5[:]

	sign, err := rsa.SignPKCS1v15(rand.Reader, prvKey.(*rsa.PrivateKey), crypto.MD5, hashed)
	if err != nil {
		return "", err
	}

	signString := base64.StdEncoding.EncodeToString(sign)
	return signString, nil
}

//验签
func RsaVerifySignWithMd5(originalData, signData, publicKey string) bool {
	sign, err := base64.StdEncoding.DecodeString(signData)
	if err != nil {
		return false
	}

	key := Base64URLDecode(publicKey)

	pubKey, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		return false
	}

	pub, err := x509.ParsePKIXPublicKey(pubKey)
	if err != nil {
		return false
	}

	hashMd5 := md5.Sum([]byte(originalData))
	hashed := hashMd5[:]

	err = rsa.VerifyPKCS1v15(pub.(*rsa.PublicKey), crypto.MD5, hashed, sign)
	if err != nil {
		return false
	}

	return true
}

func Base64URLDecode(data string) string {
	var missing = (4 - len(data)%4) % 4
	data += strings.Repeat("=", missing) //字符串长度不足4倍的位补"="
	data = strings.Replace(data, "_", "/", -1)
	data = strings.Replace(data, "-", "+", -1)
	return data
}

func Base64UrlSafeEncode(data string) string {
	safeUrl := strings.Replace(data, "/", "_", -1)
	safeUrl = strings.Replace(safeUrl, "+", "-", -1)
	safeUrl = strings.Replace(safeUrl, "=", "", -1)
	return safeUrl
}
