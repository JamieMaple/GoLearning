package main

import (
	"crypto/sha256"
	"fmt"
)

var pc [8]byte

func init() {
	for i := uint(0); i < 8; i++ {
		pc[i] = byte(1 << i)
	}
}

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Println(diffHash(&c1, &c2))
}

func diffHash(hash1, hash2 *[32]byte) int {
	count := 0

	for i := 0; i < 32; i++ {
		for j := 0; j < 8; j++ {
			if hash1[i]&pc[j] != hash2[i]&pc[j] {
				count++
			}
		}
	}

	return count
}
