package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	b0 := []byte("我是谁")
	reverseUTF8(b0)
	fmt.Printf("%s", b0)
}

func reverse(b []byte) {
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
}

func reverseUTF8(b []byte) {
	for i := 0; i < len(b); {
		_, size := utf8.DecodeRune(b[i:])
		reverse(b[i : i+size])
		i += size
	}
	reverse(b)
}
