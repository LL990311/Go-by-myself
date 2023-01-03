package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "input args length error")
		os.Exit(1)
	}
	url := os.Args[1]
	id := os.Args[2]
	Outline3(url, id)
}

func Outline3(url string, id string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintln(os.Stderr, "http get error")
		os.Exit(1)
	}
	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Fprintln(os.Stderr, "html parse error")
		os.Exit(1)
	}
	node := forEachNode(doc, id, findElement)

	if node != nil {
		fmt.Println(node.Data)
	} else {
		fmt.Println("not found id")
	}
}

func forEachNode(node *html.Node, id string, findElement func(node *html.Node, id string) bool) *html.Node {
	if findElement != nil {
		if !findElement(node, id) {
			return node
		}
	}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		if n := forEachNode(c, id, findElement); n != nil {
			return n
		}
	}
	if findElement != nil {
		if !findElement(node, id) {
			return node
		}
	}
	return nil
}

func findElement(node *html.Node, id string) bool {
	if node.Type == html.ElementNode {
		for _, a := range node.Attr {
			if a.Key == "id" && a.Val == id {
				return false
			}
		}
	}
	return true
}
