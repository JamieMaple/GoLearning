package main

import (
	"fmt"
)

func main() {
	s := []string{"hello", "hello", "s", "hello", "hello"}
	s = removeSame(s)
	fmt.Println(s)
}

func removeSame(s []string) []string {
	count := 0
	for i := range s {
		for j := i + 1; j < len(s)-count; j++ {
			if s[i] == s[j] {
				count++
				copy(s[j:], s[j+1:])
				j--
			}
		}
	}
	return s[:len(s)-count]
}
