package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	for _, link := range os.Args[1:] {
		resp, err := http.Get(link)
		if err != nil {
			fmt.Fprintf(os.Stderr, "err fetch %s : %v\n", link, err)
			os.Exit(1)
		}
		doc, err := html.Parse(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "err parse: %v\n", err)
			os.Exit(1)
		}
		node := getElementById(doc, "wrapper")
		fmt.Println(node)
	}
}

func getElementById(doc *html.Node, id string) *html.Node {
	return forEach(doc, isElement, id)
}

func forEach(n *html.Node, pre func(*html.Node, string) bool, id string) *html.Node {
	if pre != nil {
		if pre(n, id) {
			return n
		}
	}

	// foreach
	if n.FirstChild != nil {
		node := forEach(n.FirstChild, pre, id)
		if node != nil {
			return node
		}
	}
	if n.NextSibling != nil {
		node := forEach(n.NextSibling, pre, id)
		if node != nil {
			return node
		}
	}

	return nil
}

func isElement(n *html.Node, id string) bool {
	if n.Type == html.ElementNode {
		for _, attribute := range n.Attr {
			if attribute.Key == "id" {
				return attribute.Val == id
			}
		}
	}
	return false
}

