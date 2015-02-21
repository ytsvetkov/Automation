package Automation

type NFA struct {
	states  Set
	rules   NRuleBook
	accept  Set
	start   Set
	current Set
	reject  string
}

//Return new Non-deterministic Finite Automata
func NewNFA(start, states, accept Set, reject string, rules NRuleBook) *NFA {
	return &NFA{start: start, current: start, states: states, accept: accept, reject: reject, rules: rules}
}

func (n *NFA) String() string {
	return "Current state: " + n.current.String()
}

//Whether the string so far is part of the language.
func (d *NFA) Accepting() bool {
	return d.accept.Intersection(d.current).Cardinality() != 0
}

//Returns the machine back in the 'start' state.
func (n *NFA) Restart() {
	n.current = n.start
}

//Whether the string so far is not part of the language.
func (d *NFA) Rejecting() bool {
	return d.current.Contains(d.reject)
}

//Reads a single character at a time and process it.
func (n *NFA) ReadCharacter(char string) {
	if n.Rejecting() {
		n.current.Add(n.reject)
		return
	}

	a := n.current.Values()
	n.current = NewSet()
	for _, member := range a {
		b := n.rules.GetRuleEnd(member, char)
		n.current.AddSet(b)
	}

	if n.current.Cardinality() == 0 {
		n.current.Add(n.reject)
	}
}

//Reads strings and process it one character at a time.
func (n *NFA) ReadString(word string) {
	for _, char := range word {
		if n.Rejecting() {
			return
		}
		n.ReadCharacter(string(char))
	}
}

//Returns the set of all states that
//the machine has.
func (n *NFA) GetAllStates() Set {
	return n.states
}

//Returns the set of states of all which are reachable
//from the given state with the given charackter.
func (n *NFA) GetAllTransitionsFor(from, with string) (set Set) {
	if tran, ok := n.rules[from]; ok != false {
		if to, okk := tran[with]; okk != false {
			set.AddSet(to)
			return
		}
	}
	return
}

//Returns all reachable states of the machine.
func (n *NFA) GetAllEnds() (set Set) {
	for _, b := range n.rules {
		for _, j := range b {
			set.AddSet(j)
		}
	}
	return
}

//Return the set of starting states
func (n *NFA) GetStartStates() Set {
	return n.start
}

//Return the set of all accept states
func (n *NFA) GetAcceptStates() Set {
	return n.accept
}

//Return the set of all reject states
func (n *NFA) GetReject() string {
	return n.reject
}

//Returns all rules.
func (n *NFA) GetAllRules() []*Rule {
	return n.rules.GetAllRules()
}

//Returns the alphabet with wich the
//machine is working with.
func (n *NFA) GetAlphabet() Set {
	letters := NewSet()
	for _, rest := range n.rules {
		for letter := range rest {
			letters.Add(letter)
		}
	}
	return letters
}

//Returns all transition pairs from the
//given state.
func (n *NFA) GetFromState(from string) [][2]string {
	return n.rules.GetFromState(from)
}

////Return the current state of the machine
func (n *NFA) GetCurrentState() Set {
	return n.current
}
