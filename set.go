package Automation

type Set map[string]struct{}

//Return new empty set
func NewSet() Set {
	return make(Set)
}

//Returns slice, initialised with the
//strings in the set.
func SetFromSlice(words []string) Set {
	set := NewSet()
	for _, word := range words {
		set.Add(word)
	}
	return set
}

//Adds a single string to the set.
func (s Set) Add(word string) {
	s[word] = struct{}{}
}

//Adds all the elements of the given set
//to the current one.
func (s Set) AddSet(set Set) {
	for member, _ := range set {
		s[member] = struct{}{}
	}
}

//Reterns whether the given string is
//in the set.
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

//Returns slice with the strings
//in the set.
func (s Set) Values() []string {
	val := make([]string, 0, len(s))
	for member, _ := range s {
		val = append(val, member)
	}
	return val
}

//Return the size of the set i.e.
//the number of elements in the set.
func (s Set) Cardinality() int {
	return len(s)
}

//Return a new set, which is the interesction
//of the given two i.e. all elements which
//belong to both sets.
func (s Set) Intersection(other Set) Set {
	set := NewSet()
	if s.Cardinality() <= other.Cardinality() {
		for member, _ := range s {
			if other.Contains(member) {
				set.Add(member)
			}
		}
	} else {
		for member, _ := range other {
			if s.Contains(member) {
				set.Add(member)
			}
		}
	}
	return set
}

//Return a new set, which is the union
//of the given two i.e. all in both sets
func (s Set) Union(other Set) Set {
	set := NewSet()
	for member, _ := range s {
		set.Add(member)
	}
	for member, _ := range other {
		set.Add(member)
	}
	return set
}

//Checks whether the sets have the same elements.
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
