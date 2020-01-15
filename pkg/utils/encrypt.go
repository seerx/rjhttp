package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

//MD5 MD5 编码
func MD5(val string) string {
	data := []byte(val)
	has := md5.Sum(data)
	return fmt.Sprintf("%x", has)
}

// MD5s MD5 加密多个字符串，以 : 分割
func MD5s(strs ...string) string {
	str := ""
	for _, val := range strs {
		if str == "" {
			str = val
		} else {
			str += ":" + val
		}
	}
	return MD5(str)
}

// MD5f 格式化后再进行 MD5 编码
func MD5f(format string, args ...interface{}) string {
	return MD5(fmt.Sprintf(format, args))
}

// AESEncrypt AES 加密
func AESEncrypt(src, k, n string) (string, error) {
	// The key argument should be the AES key, either 16 or 32 bytes
	// to select AES-128 or AES-256.
	key := []byte(k)
	plaintext := []byte(src)

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	nonce, _ := hex.DecodeString(n)

	aesGcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	cipherText := aesGcm.Seal(nil, nonce, plaintext, nil)

	return fmt.Sprintf("%x", cipherText), nil
}

// AESDecrypt AES 解密
func AESDecrypt(src, k, n string) (string, error) {
	// The key argument should be the AES key, either 16 or 32 bytes
	// to select AES-128 or AES-256.
	key := []byte(k)
	cipherText, _ := hex.DecodeString(src)

	nonce, _ := hex.DecodeString(n)

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	aesGcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	plainText, err := aesGcm.Open(nil, nonce, cipherText, nil)
	if err != nil {
		return "", err
	}

	return string(plainText), nil
}
