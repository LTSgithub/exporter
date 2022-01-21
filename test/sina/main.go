package main

import (
	"fmt"
	"hash/crc32"
)

func String(s string) uint32 {
	return crc32.ChecksumIEEE([]byte(s))
}

func main() {
	s := "sha1 this string"

	fmt.Println(String(s))

}
