package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"hash"
	"strings"
)

// HmacSum Hash-based Message Authentication Code
func HmacSum(key string, hashFunc func() hash.Hash, elems ...string) []byte {
	h := hmac.New(hashFunc, []byte(key))
	h.Write([]byte(strings.Join(elems, "")))
	fmt.Println(hex.EncodeToString(h.Sum([]byte(""))))
	return h.Sum([]byte(""))
}

// ValidHmacSignature 验证签名
func ValidHmacSignature(key, sign string, hashFunc func() hash.Hash, elems ...string) (bool, error) {
	mac, err := hex.DecodeString(sign)
	if err != nil {
		return false, err
	}
	return hmac.Equal(mac, HmacSum(key, hashFunc, elems...)), nil
}

func Signature(accessKeySecret, method, uri, body string) string {
	hashFunc := sha256.New // 默认使用 sha256 作为 hash 函数
	h := hmac.New(hashFunc, []byte(accessKeySecret))
	h.Write([]byte(method + uri + body))
	return hex.EncodeToString(h.Sum(nil))
}

func main() {
	test := "test"
	fmt.Println(Signature("WXpCbE5Ua3pObUV5TmpReU5HRTFabUZqTUdZNFpUQmxObUUzWVRrMk1HTQ==", "POST", "api/v1/cloudenv/account", test))
	print(ValidHmacSignature("WXpCbE5Ua3pObUV5TmpReU5HRTFabUZqTUdZNFpUQmxObUUzWVRrMk1HTQ==",
		"e335361fa1a6adda7ecb86160a8346f04c57194e225361286d7ae3cb59080b22", sha256.New, "POST", "api/v1/cloudenv/account", test))
}
