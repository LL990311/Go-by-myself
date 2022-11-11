package main

import (
	"bytes"
	"fmt"
)

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3] + "," + s[n-3:])
}

//e3.10 non-recursion version of comma
func commaNonRecursion(s string) string {
	if len(s) <= 3 {
		return s
	}
	var buf bytes.Buffer

	buf.WriteByte('[')
	buf.WriteByte(s[0])
	for i := 1; i < len(s); i++ {
		if (len(s)-i)%3 == 0 {
			buf.WriteByte(',')
		}
		buf.WriteByte(s[i])
	}
	buf.WriteByte(']')
	return buf.String()
}

//e3.11 support signed float number
func commaAdv(s string) string {
	decimalStr := ""
	var buf bytes.Buffer
	// ignore and store the decimal part
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			decimalStr = s[i:]
			s = s[:i]
			break
		}
	}
	if len(s) <= 3 {
		return s
	}
	//check the input is signed or not
	if s[0] == '+' || s[0] == '-' {
		buf.WriteByte(s[0])
		s = s[1:]
	}
	buf.WriteByte(s[0])
	for i := 1; i < len(s); i++ {
		if (len(s)-i)%3 == 0 {
			buf.WriteByte(',')
		}
		buf.WriteByte(s[i])
	}
	return buf.String() + decimalStr
}

//e3.12
func e312(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	cur := make(map[rune]int, len(s1))
	for _, v := range s1 {
		cur[v]++
	}
	for _, v := range s2 {
		if cur[v] == 0 {
			return false
		}
		cur[v]--
	}
	return true
}

func main() {
	//s := "-12345679012.1233"
	//fmt.Println(commaNonRecursion(s))
	//fmt.Println(commaAdv(s))

	s1, s2 := "你是谁", "谁是你"
	fmt.Println(e312(s1, s2))
}
