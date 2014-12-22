package Automation

type NFA struct {
	start   *Set
	current *Set
	accept  *Set
	reject  string
	rules   NonDeterministicRuleBook
}

func NewNFA(starting []string, reject string, accept []string, rules NonDeterministicRuleBook) *NFA {
	x := NewSet()
	x.AddSlice(starting)
	accepting := NewSet()
	accepting.AddSlice(accept)
	return &NFA{start: x, current: x, accept: accepting, reject: reject, rules: rules}
}

func (n *NFA) Accepting() bool {
	intersection := n.current.Intersection(n.accept)
	return (&intersection).Cardinality() != 0
}

func (n *NFA) Rejecting() bool {
	return n.current.Contains(n.reject)
}

func (n *NFA) ReadCharacter(char string) {
	if n.current.Contains(n.reject) {
		return
	}

	current := n.current.Values()
	s := NewSet()

	for _, from := range current {
		for to, _ := range n.rules.GetRulesEnd(from, char) {
			s.Add(to)
		}
	}

	if s.Cardinality() == 0 {
		s.Add(n.reject)
	}

	n.current = s
}

func (n *NFA) ReadString(word string) {
	for _, i := range word {
		n.ReadCharacter(string(i))
	}
}

func (n *NFA) String() string {
	return "Start: " + n.start.String() + "\nCurrent: " + n.current.String() + "\nAccept: " + n.accept.String() + "\nRules: " + n.rules.String()
}
