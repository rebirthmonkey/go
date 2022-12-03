package util

import (
	"crypto/md5"
	"fmt"
)

// MD5 convert to md5 string
func MD5(in string) string {
	data := []byte(in)
	has := md5.Sum(data)
	out := fmt.Sprintf("%x", has)
	return out
}
