package main

import "fmt"

func sum(s []int, c chan int) {
	total := 0
	for _, v := range s {
		total += v
	}
	c <- total
}

// By default, sends and receives block until the other side is ready.

func main() {
	s := []int{12,321,32,3412,345,2368}
	ch := make(chan int)
	go sum(s[:len(s)/2], ch)
	go sum(s[len(s)/2:], ch)
	x, y := <-ch, <-ch

	fmt.Println(x, y, x + y)
}
