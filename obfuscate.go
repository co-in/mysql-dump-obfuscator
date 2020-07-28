package main

import (
	"crypto/md5"
	"fmt"
)

func obfuscateMD5(data string) string {
	hash := md5.Sum([]byte(data))

	return fmt.Sprintf("%x", hash)
}

func obfuscateEmail(data string) string {
	return obfuscateMD5(data)[:16] + "@google.com"
}

var obfuscate = map[string]map[string]func(data string) string{
	"user": {
		"username": obfuscateMD5,
		"auth_key": obfuscateMD5,
		"email":    obfuscateEmail,
	},
}
