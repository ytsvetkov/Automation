package Automation

type NFA struct {
	states  Set
	rules   NRuleBook
	accept  Set
	start   Set
	current Set
	reject  string
}

func NewNFA(start, states, accept Set, reject string, rules NRuleBook) *NFA {
	return &NFA{start: start, current: start, states: states, accept: accept, reject: reject, rules: rules}
}

func (n *NFA) String() string {
	return "Current state: " + n.current.String()
}

func (d *NFA) Accepting() bool {
	return !(d.accept.Intersection(d.current).Cardinality() != 0)
}

func (d *NFA) Rejecting() bool {
	return d.accept.Contains(d.reject)
}

func (n *NFA) ReadCharacter(char string) {
	if n.Rejecting() {
		return
	}
	a := n.current.Values()
	n.current = NewSet()
	for _, member := range a {
		b := n.rules.GetRuleEnd(member, char)
		n.current.AddSet(b)
	}
}

func (n *NFA) ReadString(word string) {
	for _, char := range word {
		if n.Rejecting() {
			return
		}
		n.ReadCharacter(string(char))
	}
}

func (n *NFA) GetAllStates() Set {
	return n.states
}

func (n *NFA) GetAllTransitionsFor(from, with string) (set Set) {
	if tran, ok := n.rules[from]; ok != false {
		if to, okk := tran[with]; okk != false {
			set.AddSet(to)
			return
		}
	}
	return
}

func (n *NFA) GetAllEnds() (set Set) {
	for _, b := range n.rules {
		for _, j := range b {
			set.AddSet(j)
		}
	}
	return
}

func (n *NFA) GetStartStates() Set {
	return n.start
}

func (n *NFA) GetAcceptStates() Set {
	return n.accept
}

func (n *NFA) GetReject() string {
	return n.reject
}

func (n *NFA) GetAllRules() []*Rule {
	return n.rules.GetAllRules()
}

func (n *NFA) GetAlphabet() Set {
	letters := NewSet()
	for _, rest := range n.rules {
		for letter := range rest {
			letters.Add(letter)
		}
	}
	return letters
}

func (n *NFA) GetFromState(from string) [][2]string {
	return n.rules.GetFromState(from)
}

func (n *NFA) GetCurrentState() Set {
	return n.current
}
