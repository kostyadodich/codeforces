package main

func main() {

}

type Stack interface {
	push(c rune)
	pop() rune
	isEmpty() bool
}
type stack struct {
	sliceSymbols []rune
}

func (s *stack) push(c rune) {
	s.sliceSymbols = append(s.sliceSymbols, c)
}

func (s *stack) pop() rune {
	result := s.sliceSymbols[len(s.sliceSymbols)-1]
	s.sliceSymbols = s.sliceSymbols[:len(s.sliceSymbols)-1]
	return result
}

func NewStack() Stack {
	return &stack{}
}

func (s *stack) isEmpty() bool {
	return len(s.sliceSymbols) == 0
}

func isValid(s string) bool {
	symbols := map[rune]rune{
		'[': ']',
		'(': ')',
		'{': '}',
	}
	st := NewStack()

	for _, c := range s {
		if isOpening(c) {
			st.push(c)
		} else {
			if st.isEmpty() {
				return false
			}
			if symbols[st.pop()] != c {
				return false
			}
		}
	}

	return st.isEmpty()
}

func isOpening(c rune) bool {
	return c == '{' || c == '(' || c == '['
}
