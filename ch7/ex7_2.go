package main

import (
	"fmt"
	"io"
	"io/ioutil"
)

func main() {
	w, c := CountingWriter(ioutil.Discard)
	fmt.Fprintf(w, "Hello, world!\n")
	fmt.Println(*c)
}

type ByteWriter struct {
	w    io.Writer
	bits int64
}

func (b *ByteWriter) Write(p []byte) (int, error) {
	n, err := b.w.Write(p)
	b.bits += int64(n)
	return n, err
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	c := ByteWriter{w, 0}
	return &c, &c.bits
}
