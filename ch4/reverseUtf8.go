package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	b := []byte("哈哈哈 嘿嘿嘿 嚯嚯嚯")
	revUTF8(b)
	fmt.Printf("%s\n", b)
}

func revUTF8(b []byte) {
	for i := 0; i < len(b); {
		fmt.Println(b)
		_, size := utf8.DecodeRune(b[i:])
		rev(b[i : i+size])
		fmt.Println(b)
		fmt.Println("----------------------------------")
		i += size
	}
	rev(b)
	fmt.Println(b)
}

func rev(b []byte) {
	last := len(b) - 1
	for i := 0; i < len(b)/2; i++ {
		b[i], b[last-i] = b[last-i], b[i]
	}
}
