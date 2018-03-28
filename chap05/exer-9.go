package main

import (
	"fmt"
	"strings"
)

func main() {
	s := expand("s$foodfsoo$foofo$foo", f)
	fmt.Println(s)
}

func f(sub string) string {
	return "bar"
}

func expand(s string, f func(string) string) string {
	const subStr = "$foo"

	return strings.Replace(s, subStr, f("foo"), -1)
}

