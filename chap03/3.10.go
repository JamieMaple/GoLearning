package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(comma("12"))
	fmt.Println(comma("12323"))
	fmt.Println(comma("1234567"))
}

func comma(s string) string {
	if len(s) <= 3 {
		return s
	}
	var str string
	for i := len(s); i > 0; i -= 3 {
		if str != "" {
			str = "," + str
		}
		str = s[int(math.Max(float64(i-3), 0)):i] + str
	}
	return str
}
