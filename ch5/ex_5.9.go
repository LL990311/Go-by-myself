package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "foofoofoo"
	fmt.Println(expand(s, replace))

}

func expand(s string, f func(string) string) string {
	return f(s)
}

func replace(s string) string {
	if strings.Contains(s, "foo") {
		s = strings.Replace(s, "foo", "$foo", -1)
	}
	return s
}
