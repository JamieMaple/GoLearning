package main

import (
	"flag"
	"fmt"
)

var ok *int
var s string

func main() {
	ok = flag.Int("ok", 233, "ok - test")
	flag.StringVar(&s, "name", "default", "name - test")
	flag.Parse()
	fmt.Println(*ok, s)
}
