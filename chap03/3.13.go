package main

import "fmt"

func main() {
	const (
		KB = 1024 << (10 * iota)
		MB
		GB
		TB
		PB
		EB
		ZB
		YB
	)
	fmt.Println(KB, MB, GB, TB, PB, EB, ZB, YB)
}
