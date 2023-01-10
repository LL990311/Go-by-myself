package main // ex 5.16

import (
	"fmt"
	"strings"
)

func stringJoin(vals ...string) string {
	if len(vals) < 2 {
		return ""
	} else if len(vals) == 1 {
		return vals[0]
	}

	sep := vals[len(vals)-1]
	curString := vals[:len(vals)-1]

	n := len(vals) * len(vals[len(vals)-1])
	for i := 0; i < len(vals); i++ {
		n += len(vals[i])
	}

	var b strings.Builder
	b.Grow(n)
	b.WriteString(vals[0])
	for _, s := range curString[1:] {
		b.WriteString(sep)
		b.WriteString(s)
	}
	return b.String()
}

func main() {
	fmt.Println(stringJoin("aaa", "bbb", "ccc", "~"))
}
