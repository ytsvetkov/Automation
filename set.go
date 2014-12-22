package Automation

type Set map[string]struct{}

func NewSet() *Set {
	x := make(Set)
	return &x
}

func (s *Set) Cardinality() int {
	return len(*s)
}

func (s *Set) Add(value string) {
	(*s)[value] = struct{}{}
}

func (s *Set) AddSlice(values []string) {
	for _, value := range values {
		s.Add(value)
	}
}

func (s *Set) Contains(value string) bool {
	_, ok := (*s)[value]
	return ok
}

func (s *Set) Remove(value string) {
	delete(*s, value)
}

func (s *Set) Clear() {
	s = NewSet()
}

func (s *Set) Values() []string {
	values := make([]string, 0, len(*s))
	for value, _ := range *s {
		values = append(values, value)
	}
	return values
}

func (s *Set) String() string {
	values := s.Values()
	str := ""
	for _, word := range values {
		str += word + ", "
	}
	return "{" + str + "}"
}

func (s *Set) Intersection(other *Set) Set {
	intesection := NewSet()
	if s.Cardinality() < other.Cardinality() {
		for _, elem := range s.Values() {
			if other.Contains(elem) {
				intesection.Add(elem)
			}
		}
	} else {
		for _, elem := range other.Values() {
			if s.Contains(elem) {
				intesection.Add(elem)
			}
		}
	}
	return *intesection
}
