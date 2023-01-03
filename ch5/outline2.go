package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "please input url")
	}
	url := os.Args[1]
	Outline2(url)
}

func Outline2(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintln(os.Stderr, "http get error happen")
	}
	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Parse error happen")
	}
	outline2(doc, startElement, endElement)
}

func outline2(node *html.Node, start, end func(node *html.Node)) {
	if start != nil {
		start(node)
	}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		outline2(c, start, end)
	}
	if end != nil {
		end(node)
	}
}

var depth int

//func startElement(n *html.Node) {
//	if n.Type == html.ElementNode {
//		fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
//		depth++
//	}
//}
//
//func endElement(n *html.Node) {
//	if n.Type == html.ElementNode {
//		depth--
//		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
//	}
//}

//ex 5.7
func startElement(n *html.Node) {
	switch n.Type {
	case html.ElementNode:
		fmt.Printf("%*s<%s", depth*2, "", n.Data)
		for _, a := range n.Attr {
			fmt.Printf(" %s=%q", a.Key, a.Val)
		}
		if n.Data == "img" && n.FirstChild == nil {
			fmt.Println("/>")
		} else {
			fmt.Println(">")
		}
		depth++
	case html.CommentNode:
		fmt.Printf("%*s<!--%s-->\n", depth*2, "", n.Data)
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		if n.Data != "img" || n.FirstChild != nil {
			fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
		}
	}
}
