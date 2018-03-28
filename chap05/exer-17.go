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
			fmt.Fprint(os.Stderr, "fetch %s err:  %v\n", link, err)
			os.Exit(1)
		}
		doc, err := html.Parse(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprint(os.Stderr, "parse err: %v\n", err)
			os.Exit(1)
		}
		nodes := ElementsByTagNames(doc, "div", "h1", "h2")
		for _, node := range nodes {
			fmt.Printf("%v\n", node)
		}
	}
}

func contains(arr []string, s string) bool {
	for _, temp := range arr {
		if temp == s {
			return true
		}
	}
	return false
}

func ElementsByTagNames(n *html.Node, names ...string) []*html.Node {
	var nodes []*html.Node
	pre := func(n *html.Node) bool {
		if n.Type == html.ElementNode && contains(names, n.Data) {
			nodes = append(nodes, n)
			return true
		} else {
			return false
		}
	}

	forEachNode(n, pre, nil)

	return nodes
}

func forEachNode(n *html.Node, pre, post func(n *html.Node) bool) *html.Node {
	if pre != nil {
		pre(n)
	}

	if n.FirstChild != nil {
		forEachNode(n.FirstChild, pre, post)
	}

	if post != nil {
		post(n)
	}

	if n.NextSibling != nil {
		forEachNode(n.NextSibling, pre, post)
	}

	return nil
}

