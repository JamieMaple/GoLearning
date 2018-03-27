package main

import (
	"fmt"
	"net/http"
	"os"
	"unicode/utf8"

	"golang.org/x/net/html"
)

func main() {
	for _, link := range os.Args[1:] {
		words, images, err := CountWordsAndImages(link)
		if err != nil {
			fmt.Fprintf(os.Stderr, "err: %v", err)
			break
		}
		fmt.Printf("link: %s\n\twords: %d\n\timages: %d\n", link, words, images)
	}
}

func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
	}
	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(n *html.Node) (words, images int) {
	if n.Type == html.ElementNode && n.Data == "img" {
		images = 1
	}
	if n.Type == html.TextNode {
		ParentData := n.Parent.Data

		if ParentData != "script" && ParentData != "style" && ParentData != "noscript" {
			words = countWords(n.Data)
		}
	}

	if n.FirstChild != nil {
		wordsIncrement, imagesIncrement := countWordsAndImages(n.FirstChild)
		words += wordsIncrement
		images += imagesIncrement
	}
	if n.NextSibling != nil {
		wordsIncrement, imagesIncrement := countWordsAndImages(n.NextSibling)
		words += wordsIncrement
		images += imagesIncrement
	}

	return
}

func countWords(line string) (count int) {
	for i := 0; i < len(line); {
		_, size := utf8.DecodeRuneInString(line[i:])
		count++
		i += size
	}

	return
}

