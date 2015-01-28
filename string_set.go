package Automation

type Set map[string]struct{}

func NewSet() Set {
	return make(Set)
}

func SetFromSlice(words []string) Set {
	set := NewSet()
	for _, word := range words {
		set.Add(word)
	}
	return set
}

func (s Set) Add(word string) {
	s[word] = struct{}{}
}

func (s Set) AddSet(set Set) {
	for member, _ := range set {
		s[member] = struct{}{}
	}
}

func (s Set) Contains(word string) bool {
	_, ok := s[word]
	return ok
}

func (s Set) String() string {
	str := "{"
	for member, _ := range s {
		str += member + ", "
	}
	return str + "}"
}

func (s Set) Values() []string {
	val := make([]string, 0, len(s))
	for member, _ := range s {
		val = append(val, member)
	}
	return val
}

func (s Set) Cardinality() int {
	return len(s)
}

func (s Set) Intersection(other Set) Set {
	set := NewSet()
	for member, _ := range s {
		set.Add(member)
	}
	for member, _ := range other {
		set.Add(member)
	}
	return set
}

func (s Set) Clear() {
	s = NewSet()
}

func (s Set) Eq(other Set) bool {
	if s.Cardinality() != s.Cardinality() {
		return false
	}

	var flag bool
	for _, i := range s.Values() {
		flag = false
		for _, j := range other.Values() {
			if i == j {
				flag = true
			}
		}
		if !flag {
			return false
		}
	}

	for _, i := range other.Values() {
		flag = false
		for _, j := range s.Values() {
			if i == j {
				flag = true
			}
		}
		if !flag {
			return false
		}
	}

	return true
}
