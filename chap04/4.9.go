package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	words := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	input.Split(bufio.ScanWords)
	for input.Scan() {
		if input.Text() == "end" {
			break
		}
		words[input.Text()]++
	}
	if input.Err() != nil {
		fmt.Fprintln(os.Stderr, "input error: %v", input.Err())
		os.Exit(1)
	}
	for s, n := range words {
		fmt.Printf("%s\t\t%d\n", s, n)
	}
}
