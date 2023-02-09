package main

import (
	"bufio"
	"bytes"
	"fmt"
)

type WordCounter int
type LineCounter int

var wc WordCounter
var lc LineCounter = 1

func main() {
	s := "hello world!/n hello hello"
	fmt.Fprintf(&wc, s)
	fmt.Println(wc)

	fmt.Fprintln(&lc, s)
	fmt.Println(lc)
}

// write "s" into "w" return a int (nums of words)
func (w *WordCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		*w++
	}
	return len(p), nil
}

func (l *LineCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		*l++
	}
	return len(p), nil
}
