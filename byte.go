package main

import (
	"unsafe"
)

// Bytes2String 字节集转字符串
func Bytes2String(b []byte) string {
	return unsafe.String(unsafe.SliceData(b), len(b))
}
