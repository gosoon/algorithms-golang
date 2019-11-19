package stack

import "fmt"

// ArrayStack is define a stack
type ArrayStack struct {
	data   []interface{}
	top    int
	length int
}

// NewArrayStack is new a array stack
func NewArrayStack(c int) *ArrayStack {
	return &ArrayStack{
		data:   make([]interface{}, c),
		top:    -1,
		length: c,
	}
}

// Push is push a element to stack top
func (s *ArrayStack) Push(v interface{}) bool {
	if !s.IsFull() {
		s.top++
		s.data[s.top] = v
		s.length++
		return true
	}
	return false
}

// Pop is pop a element
func (s *ArrayStack) Pop() (interface{}, bool) {
	if !s.IsEmpty() {
		v := s.data[s.top]
		s.top--
		s.length--
		return v, true
	}
	return 0, false
}

// IsFull is valid stack is full
func (s *ArrayStack) IsFull() bool {
	if s.top == s.length {
		return true
	}
	return false
}

// IsEmpty is valid stack is empty
func (s *ArrayStack) IsEmpty() bool {
	if s.top < 0 {
		return true
	}
	return false
}

// Search is find a val in stack
func (s *ArrayStack) Search(v interface{}) bool {
	for i := 0; i <= s.top; i++ {
		if s.data[i] == v {
			return true
		}
	}
	return false
}

// Print is print all data int stack
func (s *ArrayStack) Print(info string) {
	fmt.Println(info)
	for i := 0; i <= s.top; i++ {
		fmt.Println(s.data[i])
	}
}
