package main

import (
	"bytes"
	"fmt"
)

type IntSet struct {
	words []uint64
}

func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

//ex 6.1
func (s *IntSet) Len() int {
	cnt := 0
	for _, word := range s.words {
		for mask := 0; mask < 64; mask++ {
			if word&(1<<mask) != 0 {
				cnt++
			}
		}
	}
	return cnt
}

//ex 6.1
func (s *IntSet) Remove(x int) {
	word, bit := x/64, x%64
	if word > len(s.words) {
		return
	}
	s.words[word] &^= 1 << bit
}

//ex 6.1
func (s *IntSet) Clear() {
	s.words = nil
}

//ex 6.1
func (s *IntSet) Copy() *IntSet {
	n := &IntSet{}
	n.words = make([]uint64, len(s.words))
	copy(n.words, s.words)
	return n
}

//ex 6.2
func (s *IntSet) AddAll(n ...int) {
	for _, x := range n {
		word, bit := x/64, x%64
		for word >= len(s.words) {
			s.words = append(s.words, 0)
		}
		s.words[word] |= 1 << bit
	}
}

/*
 10110
 11000
 11110
*/
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

//ex 6.3
/*
 10111
 00101
 00101
*/
func (s *IntSet) IntersectionWith(t *IntSet) {
	minLen := _min(len(s.words), len(t.words))
	for i := 0; i < minLen; i++ {
		s.words[i] &= t.words[i]
	}
	s.words = s.words[:minLen]
}

func _min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

/*
 11101
 10011
 01100
*/
func (s *IntSet) DifferentiateWith(t *IntSet) {
	for i := 0; i < len(s.words); i++ {
		if i < len(t.words) {
			temp := s.words[i]
			s.words[i] ^= t.words[i]
			s.words[i] &= temp
		} else {
			return
		}
	}
}

/*
 11111
 01001
 10110
*/
func (s *IntSet) SymDifferentiateWith(t *IntSet) {
	sLen := len(s.words)
	var i = 0
	for ; i < sLen; i++ {
		if i < len(t.words) {
			s.words[i] ^= t.words[i]
		} else {
			return
		}
	}
	for ; i < len(t.words); i++ {
		s.words = append(s.words, t.words[i])
	}

}

func _max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func main() {
	var x, y IntSet
	x.AddAll(1, 4, 3)
	y.AddAll(1, 2)

	fmt.Println(x.String())
	fmt.Println(y.String())

	x.UnionWith(&y)

	fmt.Println(x.String())

}
