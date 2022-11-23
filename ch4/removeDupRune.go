// ex4.6
package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func removeDupRune(arr []byte) []byte {
	for i := 0; i < len(arr); {
		first, size := utf8.DecodeRune(arr[i:])
		if unicode.IsSpace(first) {
			second, _ := utf8.DecodeRune(arr[i+size:])
			if unicode.IsSpace(second) {
				copy(arr[i:], arr[i+size:])
				arr = arr[:len(arr)-size]
			}
		}
		i += size
	}
	return arr
}

func main() {
	arr := []byte("哈哈哈哈    哈哈  嘿嘿 哈哈")
	arr = removeDupRune(arr)
	fmt.Printf("%s\n", arr)
}
