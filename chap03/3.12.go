package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(cmpStrings("sd我", "sd"))
}

func cmpStrings(s1, s2 string) bool {
	if len(s1) != len(s2) || s1 == s2 {
		return false
	}
	r := []rune(s1)
	for i, length := 0, len(r); i < length; i++ {
		//fmt.Println(string(r[i]))
		wordPos := strings.Index(s2, string(r[i]))
		if wordPos < 0 || wordPos == i {
			return false
		}
	}
	return true
}
