package main

import (
	"bytes"
	"fmt"
)

const wordSize = 32 << (^uint(0) >> 63)

type IntSet struct {
	words []uint
}

func (s *IntSet) Has(x int) bool {
	word, bit := x/wordSize, uint(x%wordSize)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *IntSet) Add(x int) {
	word, bit := x/wordSize, uint(x%wordSize)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < wordSize; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", wordSize*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func (s *IntSet) Len() int {
//	todo
	return 0
}

func (s *IntSet) Remove(x int) {
// todo
}

func (s *IntSet) Clear() {
	for i := range s.words {
		s.words[i] = 0
	}
}

func (s *IntSet) Copy() *IntSet {
	var clone *IntSet

	clone.words = make([]uint, len(s.words))
	copy(clone.words, s.words)

	return clone
}

func (s *IntSet) AddAll(nums ...int) {
	for _, num := range nums {
		s.Add(num)
	}
}

func (s *IntSet) Elems() (e []int) {
	//var e []int
	for i, word := range s.words {
		for j := 0; j < wordSize; j++ {
			if word&(1<<uint(j)) != 0 {
				e = append(e, i*wordSize+j)
			}
		}
	}
	return
}

func main() {
	var a IntSet
	a.AddAll(1, 2, 3, 20)
	fmt.Println(a.Elems())
}
