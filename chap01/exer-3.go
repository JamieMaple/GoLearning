package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	for _, arg := range os.Args[1:] {
		f, err := os.Open(arg)
		if err != nil {
			fmt.Fprint(os.Stderr, "err: %v\n", err)
			continue
		}
		// get input into counts
		input := bufio.NewScanner(f)
		for input.Scan() {
			counts[input.Text()]++
		}
		for _, n := range counts {
			if n > 1 {
				fmt.Println(arg)
				break
			}
		}
	}
}
