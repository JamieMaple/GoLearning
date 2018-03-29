package main

import "fmt"

func main() {
	result := f()
	fmt.Println(result)
}

func f() (result int) {
	fmt.Println("no return!")
	defer func() {
		switch p := recover(); p {
		case "error":
			fmt.Println("error happen!")
			result = 1
		}
	}()
	panic("error")
}

