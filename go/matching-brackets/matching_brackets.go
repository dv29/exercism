package brackets

import (
	"errors"
)

// Stack stack to hold brackets
type Stack []rune

// Push pushes an open bracket on stack
func (s *Stack) Push(r rune) {
	*s = append(*s, r)
}

// Pop removes the last element from stack and returns it
func (s *Stack) Pop() (rune, error) {
	if len(*s) == 0 {
		return 0, errors.New("empty stack")
	}
	ele := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]

	return ele, nil
}

// Seek returns the last element from stack
func (s *Stack) Seek() (rune, error) {
	if len(*s) == 0 {
		return 0, errors.New("empty stack")
	}
	ele := (*s)[len(*s)-1]

	return ele, nil
}

// New initialize new stack
func New() *Stack {
	return &Stack{}
}

var bracketPair = map[rune]rune{
	'}': '{',
	']': '[',
	')': '(',
}

// Bracket checks if the pairs of bracket are valid
func Bracket(inp string) bool {
	s := New()
	for _, v := range inp {
		if v == '{' || v == '[' || v == '(' {
			s.Push(v)
		} else if v == '}' || v == ']' || v == ')' {
			x, err := s.Seek()
			if err != nil {
				return false
			}
			p := bracketPair[x]
			if x != p {
				return false
			}
			if _, err := s.Pop(); err != nil {
				return false
			}
		}
	}
	return true
}
