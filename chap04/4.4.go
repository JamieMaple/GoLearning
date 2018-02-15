package main

import (
	"fmt"
)

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	rotate(s, 3)
	fmt.Println(s)
}

func rotate(s []int, i int) bool {
	if i > len(s) {
		return false
	}
	reverse(s[:i])
	reverse(s[i:])
	reverse(s)
	return true
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
