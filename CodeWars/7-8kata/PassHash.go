package main

import (
	"crypto/md5"
	"fmt"
)

func PassHash(str string) string {
	hash := md5.Sum([]byte(str))
	return fmt.Sprintf("%x", hash)
}

func main() {
	fmt.Println(PassHash("password"))
}
