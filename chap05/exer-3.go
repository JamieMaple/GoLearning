package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error fetch test: %v", err)
	}
	getText(doc)
}

func getText(n *html.Node) {
	if n.Type == html.TextNode && n.Data != "script" && n.Data != "style" {
		fmt.Println(n.Data)
	}
	if n.FirstChild != nil {
		getText(n.FirstChild)
	}
	if n.NextSibling != nil {
		getText(n.NextSibling)
	}
}

