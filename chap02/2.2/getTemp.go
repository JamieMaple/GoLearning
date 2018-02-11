// need 2.1 package
package main

import (
	"bufio"
	"fmt"
	"github.com/jamie/goland/tempconv"
	"os"
	"strconv"
)

func main() {
	if len(os.Args[1:]) == 0 {
		var cels []string
		var input = bufio.NewScanner(os.Stdin)
		fmt.Println("请输入摄氏度，以 end 结束")
		for input.Scan() {
			if input.Text() == "end" {
				break
			} else {
				cels = append(cels, input.Text())
			}
		}
		getTemp(cels)
	} else {
		getTemp(os.Args[1:])
	}
}

func getTemp(cels []string) {
	for _, cel := range cels {
		v, err := strconv.Atoi(cel)
		if err != nil {
			continue
		}
		c := tempconv.Celsius(v)
		fmt.Println("Celsius：", c)
		fmt.Println("Fahrenheit", tempconv.CToF(c))
	}
}
