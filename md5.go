package tools

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
)

// ByteGetMd5Byte 获取md5
func ByteGetMd5Byte(bytes []byte) ([]byte, error) {
	fileMd5 := md5.New()
	_, err := fileMd5.Write(bytes)
	if err != nil {
		return nil, err
	}
	return fileMd5.Sum(nil), nil
}

// ByteGetMd5String 获取md5
func ByteGetMd5String(bytes []byte) (string, error) {
	md5, err := ByteGetMd5Byte(bytes)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(md5), nil
}

// ByteGetMd5Base64 获取md5
func ByteGetMd5Base64(bytes []byte) (string, error) {
	md5, err := ByteGetMd5Byte(bytes)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(md5), nil
}

// StringGetMd5Byte 获取md5
func StringGetMd5Byte(str string) ([]byte, error) {
	fileMd5 := md5.New()
	_, err := fileMd5.Write(String2Bytes(str))
	if err != nil {
		return nil, err
	}
	return fileMd5.Sum(nil), nil
}

// StringGetMd5String 获取md5
func StringGetMd5String(str string) (string, error) {
	md5, err := StringGetMd5Byte(str)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(md5), nil
}

// StringGetMd5Base64 获取md5
func StringGetMd5Base64(str string) (string, error) {
	md5, err := StringGetMd5Byte(str)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(md5), nil
}
