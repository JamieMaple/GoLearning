package main

import (
	"golang.org/x/net/html"
	"os"
	"fmt"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Println(os.Stderr, "find elements count: %v\n", err)
	}
	tagMap := make(map[string]int)
	visit(&tagMap, doc)
    fmt.Println(tagMap)
}

func visit(elemsMap *map[string]int, n *html.Node) {
	if n.Type == html.ElementNode {
		(*elemsMap)[n.Data]++
	}
	if n.FirstChild != nil {
		visit(elemsMap, n.FirstChild)
	}
	if n.NextSibling != nil {
		visit(elemsMap, n.NextSibling)
	}
}

