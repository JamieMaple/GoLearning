package main

import (
	"fmt"
	"os"
)

func main() {
	for index, val := range os.Args {
		fmt.Printf("%v : %v\n", index, val)
	}
}
