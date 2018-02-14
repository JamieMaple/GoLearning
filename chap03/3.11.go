package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(comma("12"))
	fmt.Println(comma("12323"))
	fmt.Println(comma("1234567"))
	fmt.Println(comma("-123456789"))
	fmt.Println(comma("+123456789"))
	fmt.Println(comma("1234.2341"))
	fmt.Println(comma("1234.2341233"))
}

func comma(s string) string {
	if s[0] == '-' || s[0] == '+' {
		return string(s[0]) + commaForPostive(s[1:])
	} else {
		return commaForPostive(s)
	}
}

func commaForPostive(s string) string {
	if len(s) <= 3 {
		return s
	}
	var str string
	var i int
	dotPos := strings.Index(s, ".")
	if dotPos > 0 {
		i = dotPos
	} else {
		i = len(s)
	}
	for i > 0 {
		if i-3 > 0 {
			str = "," + s[i-3:i] + str
		} else {
			str = s[:i] + str
		}
		i -= 3
	}
	if dotPos > 0 {
		str += "."
		i = dotPos + 1
		for length := len(s); i < length; i += 3 {
			if i+3 < length {
				str += s[i:i+3] + ","
			} else {
				str += s[i:]
			}
		}
	}
	return str
}
