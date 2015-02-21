package Automation

type DFA struct {
	states  Set
	rules   DRuleBook
	accept  Set
	start   string
	current string
	reject  string
}

// Returns new Deterministic Finite Automata
func NewDFA(start, reject string, states, accept Set, rules DRuleBook) *DFA {
	return &DFA{start: start, current: start, reject: reject, states: states, accept: accept, rules: rules}
}

func (d *DFA) String() string {
	return "Current state: " + d.current
}

// Returns the machine back in the 'start' state.
func (d *DFA) Restart() {
	d.current = d.start
}

// Whether the string so far is part of the language.
func (d *DFA) Accepting() bool {
	return d.accept.Contains(d.current)
}

// Whether the string so far is not part of the language.
func (d *DFA) Rejecting() bool {
	return d.current == d.reject
}

// Reads a single character at a time.
func (d *DFA) ReadCharacter(char string) {
	if d.Rejecting() {
		return
	}
	x := d.rules.GetFromTransition(d.current, char).Values()
	if len(x) != 1 {
		d.current = d.reject
		return
	}
	d.current = x[0]
}

// Reads string.
func (d *DFA) ReadString(word string) {
	for _, char := range word {
		if d.Rejecting() {
			return
		}
		d.ReadCharacter(string(char))
	}
}

// Returns the set of all states that
// the machine has.
func (d *DFA) GetAllStates() Set {
	return d.states
}

// Returns the set of states of all which are reachable
// from the given state with the given charackter.
// Because this is deterministic machine, there is going to be
// no more than one element in the set.
func (d *DFA) GetAllTransitionsFor(from, with string) Set {
	set := NewSet()
	if tran, ok := d.rules[from]; ok != false {
		if to, okk := tran[with]; okk != false {
			set.Add(to)
			return set
		}
	}
	return set
}

// Returns all reachable states of the machine.
func (d *DFA) GetAllEnds() Set {
	set := NewSet()
	for _, b := range d.rules {
		for _, j := range b {
			set.Add(j)
		}
	}
	return set
}

// Returns the alphabet with wich the
// machine is working with.
func (d *DFA) GetAlphabet() Set {
	letters := NewSet()
	for _, rest := range d.rules {
		for letter := range rest {
			letters.Add(letter)
		}
	}
	return letters
}

// Return the set of starting states.
// This is deterministic machine so
// there could be only one such state.
func (d *DFA) GetStartStates() Set {
	start := NewSet()
	start.Add(d.start)
	return start
}

// Return the set of all accept states
func (d *DFA) GetAcceptStates() Set {
	return d.accept
}

// Return the set of all reject states
func (d *DFA) GetReject() string {
	return d.reject
}

// Returns all rules.
func (d *DFA) GetAllRules() []*Rule {
	return d.rules.GetAllRules()
}

// Returns all transition pairs from the
// given state. Because this is deterministic
// machine, there cant be more than one element
// in the slice.
func (d *DFA) GetFromState(from string) [][2]string {
	return d.rules.GetFromState(from)
}

// // Return the current state of the machine
func (d *DFA) GetCurrentState() string {
	return d.current
}
