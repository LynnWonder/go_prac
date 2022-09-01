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

type Foo struct {
	Number int    `json:"number"`
	Title  string `json:"title"`
}

func main() {
	fmt.Println(Signature("xx", "POST", "/api/v1/test", "test"))
	print(ValidHmacSignature("xx",
		"c2322ce78f8e016ad696cfe287058776ec7271bb7643876e05374ad85c8e0740", sha256.New, "POST", "/api/v1/test", "test"))
}
