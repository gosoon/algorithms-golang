package stack

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	s := NewArrayStack(5)
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	s.Print("")

	if v, ok := s.Pop(); ok {
		fmt.Println(v)
	}
	s.Print("")

	if found := s.Search(3); found {
		fmt.Println("found")
	} else {
		fmt.Println("not found")
	}
}
