package main

type Stack []string

// IsEmpty: check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Length() int {
	return len(*s)
}

// Push a new value onto the stack
func (s *Stack) Push(str string) {
	*s = append(*s, str)
}

// return top element of stack. Return false if stack is empty
func (s *Stack) Top() (string, bool) {
	if s.IsEmpty() {
		return "", false
	}
	return (*s)[len(*s)-1], true
}

// Remove and return top element of stack. Return false if stack is empty.
func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index]      // Remove it from the stack by slicing it off.
		return element, true
	}
}
