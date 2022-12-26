package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	//for _, url := range os.Args[1:] {
	//	// e1.8 if url doesn't have http:// prefix, add it automatically
	//	if !strings.HasPrefix(url, "http://") {
	//		url = "http://" + url
	//	}
	//	resp, err := http.Get(url)
	//	if err != nil {
	//		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
	//		os.Exit(1)
	//	}
	//	//b, err := ioutil.ReadAll(resp.Body)
	//	_, err = io.Copy(os.Stdout, resp.Body) // e1.7 use isCopy to read resp.Body
	//	if err != nil {
	//		fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
	//		os.Exit(1)
	//	}
	//	// e1.9 print http status code
	//	fmt.Printf("\n%s\n", resp.Status)
	//}
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}
