package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

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

		start, end := getStartAndEndFuncs()
		forEach(doc, start, end)
	}
}

func forEach(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	// foreach
	if n.FirstChild != nil {
		forEach(n.FirstChild, pre, post)
	}
	if post != nil {
		post(n)
	}
	if n.NextSibling != nil {
		forEach(n.NextSibling, pre, post)
	}
	//for c := n.FirstChild; c != nil; c = c.NextSibling {
	//	forEach(c, pre, post)
	//}
}

var singerElements = [...]string{"img", "input", "br"}

func contains(data string) (isContain bool) {
	for _, str := range singerElements {
		if str == data {
			isContain = true
			return
		}
	}
	isContain = false
	return
}

func getAttributes(n *html.Node) map[string]string {
	var attributes = make(map[string]string)
	if n.Type == html.ElementNode {
		for _, attribute := range n.Attr {
			attributes[attribute.Key] = attribute.Val
		}
	}
	return attributes
}

func joinAttributes(attributes map[string]string) string {
	var arr []string
	for k, v := range attributes {
		arr = append(arr, fmt.Sprintf("%s=%q", k, v))
	}
	return strings.Join(arr, " ")
}

func getStartAndEndFuncs() (startElement, endElement func(*html.Node)) {
	var depth = 1

	startElement = func(n *html.Node) {
		if n.Type == html.ElementNode {
			attributes := joinAttributes(getAttributes(n))
			if attributes != "" {
				attributes = " " + attributes
			}
			switch data := n.Data; {
			case contains(data):
				fmt.Printf("%*s<%s%s/>\n", depth*2, "", n.Data, attributes)
			default:
				fmt.Printf("%*s<%s%s>\n", depth*2, "", n.Data, attributes)
			}
			depth++
		}
	}

	endElement = func(n *html.Node) {
		if n.Type == html.ElementNode {
			depth--
			switch data := n.Data; {
			case contains(data):
			default:
				fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
			}
		}
	}
	return
}

