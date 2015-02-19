package Automation

type item struct {
	value string
	next  *item
}

type Stack struct {
	top  *item
	size int
}

func (s *Stack) Len() int {
	return s.size
}

func (s *Stack) PushChar(char string) {
	s.top = &item{value: char, next: s.top}
	s.size++
}

func (s *Stack) Push(value string) {
	for _, char := range value {
		s.PushChar(string(char))
	}
}

func (s *Stack) PushString(value string) {
	s.top = &item{value: value, next: s.top}
	s.size++
}

func (s *Stack) Pop() (value string, exists bool) {
	exists = false
	if s.size > 0 {
		value, s.top = s.top.value, s.top.next
		s.size--
		exists = true
	}
	return
}

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
