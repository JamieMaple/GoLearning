package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"os"
	"strings"
)

func main() {
	// default sha256
	use := "sha256"
	if len(os.Args) > 1 {
		use = os.Args[1]
	}
	input := []byte(handleInput())
	switch use {
	case "sha384":
		fmt.Printf("%x\n", sha512.Sum384(input))
	case "sha512":
		fmt.Printf("%x\n", sha512.Sum512(input))
	default:
		fmt.Printf("%x\n", sha256.Sum256(input))
	}
}

func handleInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimRight(input, "\n")
}
