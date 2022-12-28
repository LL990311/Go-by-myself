//ex5.5
package main

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"net/http"
	"os"
	"strings"
)

func CountWordsAndImages(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	doc, err := html.Parse(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return fmt.Errorf("parsing HTML: %s", err)
	}
	countWordsAndImages(doc)
	return nil
}

var (
	words  int
	images int
)

func countWordsAndImages(n *html.Node) {
	if n.Type == html.ElementNode && n.Data == "img" {
		images++
	}
	if n.Type == html.TextNode {
		words += len(strings.Fields(n.Data))
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		countWordsAndImages(c)
	}
	return
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "usage: PROG URL")
	}
	url := os.Args[1]
	err := CountWordsAndImages(url)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("words: %d\nimages: %d\n", words, images)
}
