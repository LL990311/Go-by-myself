//ex4.9
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	freq := make(map[string]int)
	io := bufio.NewScanner(os.Stdin)
	io.Split(bufio.ScanWords)
	for io.Scan() {
		freq[io.Text()]++
	}
	for k, v := range freq {
		fmt.Printf("%s\t%d\n", k, v)
	}
}
