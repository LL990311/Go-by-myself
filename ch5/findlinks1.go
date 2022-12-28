package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
	"strings"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}

	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}

	////ex 5.2
	//freq := make(map[string]int, 0)
	//for key, value := range elementCounts(freq, doc) {
	//	fmt.Printf("%-8s:%d\n", key, value)
	//}

	////ex 5.3
	//output(doc)
}

//// iteration
//func visit(links []string, n *html.Node) []string {
//	if n.Type == html.ElementNode && n.Data == "a" {
//		for _, a := range n.Attr {
//			if a.Key == "href" {
//				links = append(links, a.Val)
//			}
//		}
//	}
//	for c := n.FirstChild; c != nil; c = c.NextSibling {
//		links = visit(links, c)
//	}
//	return links
//}

//recursion (ex5.1) (ex5.4)
func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	if n.Type == html.ElementNode && n.Data == "script" || n.Data == "style" || n.Data == "img" {
		for _, a := range n.Attr {
			if a.Key == "src" {
				links = append(links, a.Val)
			}
		}
	}
	if n.FirstChild != nil {
		links = visit(links, n.FirstChild)
	}
	if n.NextSibling != nil {
		links = visit(links, n.NextSibling)
	}
	return links
}

//ex5.2
func elementCounts(freq map[string]int, n *html.Node) map[string]int {
	if n.Type == html.ElementNode {
		freq[n.Data]++
	}
	if n.FirstChild != nil {
		elementCounts(freq, n.FirstChild)
	}
	if n.NextSibling != nil {
		elementCounts(freq, n.NextSibling)
	}
	return freq
}

//ex 5.3
func output(n *html.Node) {
	if n.Type == html.TextNode {
		s := strings.Join(strings.Fields(n.Data), " ")
		if s != "" {
			fmt.Println(s)
		}
	}
	if n.FirstChild != nil {
		output(n.FirstChild)
	}
	if n.NextSibling != nil {
		output(n.NextSibling)
	}
}
