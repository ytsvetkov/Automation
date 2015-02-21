package Automation

type item struct {
	value string
	next  *item
}

// Stack
type Stack struct {
	top  *item
	size int
}

// Returns the length( or the size) of the stacl.
func (s *Stack) Len() int {
	return s.size
}

// Pushes a single character.
func (s *Stack) PushChar(char string) {
	s.top = &item{value: char, next: s.top}
	s.size++
}

// Pushes the symbols of the given string.
func (s *Stack) Push(value string) {
	for _, char := range value {
		s.PushChar(string(char))
	}
}

// Pushed the entire string.
func (s *Stack) PushString(value string) {
	s.top = &item{value: value, next: s.top}
	s.size++
}

// Removes the top element of the stack and returns it.
func (s *Stack) Pop() (value string, exists bool) {
	exists = false
	if s.size > 0 {
		value, s.top = s.top.value, s.top.next
		s.size--
		exists = true
	}
	return
}

// Returns the value of the top element of the stack
// withouth removing it from the stack.
func (s *Stack) Peek() (value string, exists bool) {
	exists = false
	if s.size > 0 {
		value = s.top.value
		exists = true
	}
	return
}

func NewStack() *Stack {
	return new(Stack)
}

func (s *Stack) String() string {
	ptr := s.top
	str := "["

	for i := 0; i < s.size-1; i++ {
		str += ptr.value + ", "
		ptr = ptr.next
	}

	if ptr != nil {
		return str + ptr.value + "]"
	}
	return ""
}
