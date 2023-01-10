package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
)

func ElementsByTagname(doc *html.Node, name ...string) []*html.Node {
	if doc == nil {
		return nil
	}
	var nodes []*html.Node
	if doc.Type == html.ElementNode {
		for _, s := range name {
			if doc.Data == s {
				nodes = append(nodes, doc)
			}
		}
	}
	for c := doc.FirstChild; c != nil; c = c.NextSibling {
		nodes = append(nodes, ElementsByTagname(c, name...)...)
	}
	return nodes
}

func fetch(url string) *html.Node {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintln(os.Stderr, "url invalid")
	}
	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Fprintln(os.Stderr, "parse error")
	}
	return doc
}

func main() {
	url := os.Args[1]
	doc := fetch(url)
	nodes := ElementsByTagname(doc, "h1", "h2", "h3", "img")
	for _, n := range nodes {
		for _, a := range n.Attr {
			fmt.Printf("Type: %s, %s= %q\n", n.Data, a.Key, a.Val)
		}
	}
}
