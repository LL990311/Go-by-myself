// Package sha256 ex4.1
package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

func sha256Diff(b1, b2 [32]byte) int {
	count := 0
	for i := 0; i < 32; i++ {
		count += diffCount(b1[i], b2[i])
	}
	return count
}

func diffCount(b1, b2 byte) int {
	count := 0
	for i := 0; i < 8; i++ {
		mask := byte(1 << i)
		if b1&mask != b2&mask {
			count++
		}
	}
	return count
}

var f = flag.String("flag", "sha256", "flag = sha256 | sha384 | sha512")

func main() {
	//c1 := sha256.Sum256([]byte("x"))
	//c2 := sha256.Sum256([]byte("X"))
	//fmt.Println(sha256Diff(c1, c2))

	//ex4.2
	flag.Parse()
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		bytes := input.Bytes()
		switch *f {
		case "sha256":
			fmt.Printf("sha256: %x\n", sha256.Sum256(bytes))
		case "sha384":
			fmt.Printf("sha384: %x\n", sha512.Sum384(bytes))
		case "sha512":
			fmt.Printf("sha512: %x\n", sha512.Sum512(bytes))
		}
	}
}
