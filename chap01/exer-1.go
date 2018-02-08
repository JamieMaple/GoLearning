package main

import (
	"fmt"
	"os"
)

func main() {
	var step, s string
	for _, val := range os.Args {
		s += step + val
		step = " "
	}
	fmt.Println(s)
}
