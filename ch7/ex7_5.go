package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	lr := LimitReader(strings.NewReader("hello world"), 5)
	b, err := ioutil.ReadAll(lr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "err: %v", err)
	}
	fmt.Printf("%s\n", b)
}

type LimitedReader struct {
	reader io.Reader
	bits   int64
}

func (lr *LimitedReader) Read(b []byte) (n int, err error) {
	if lr.bits <= 0 {
		return 0, io.EOF
	}
	if (int64(len(b))) > lr.bits {
		b = b[:lr.bits]
	}
	n, err = lr.reader.Read(b)
	lr.bits -= int64(n)
	return
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return &LimitedReader{r, n}
}
