package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() && len(counts) < 5 {
		counts[input.Text()]++
	}
	// ignore the error instead
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n\n", n, line)
		}
	}
}
