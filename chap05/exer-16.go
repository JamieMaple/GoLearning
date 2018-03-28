package main

import "fmt"

func main() {
	var test1 = []string{""}
	var test2 = []string{"hello", ""}
	var test3 = []string{"hello", "wrold"}
	const fill = " "

	fmt.Println(Join(test1, fill))
	fmt.Println(Join(test2, fill))
	fmt.Println(Join(test3, fill))
}

func Join(a []string, fill string) (s string) {
	for i, item := range a {
		if i == 0 {
			s = a[0]
		} else {
			s += fill + item
		}
	}
	return
}

