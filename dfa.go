package Automation

type DFA struct {
	states  Set
	rules   DRuleBook
	accept  Set
	start   string
	current string
	reject  string
}

func NewDFA(start, reject string, states, accept Set, rules DRuleBook) DFA {
	return DFA{start: start, current: start, reject: reject, states: states, accept: accept, rules: rules}
}

func (d DFA) Accepting() bool {
	return d.accept.Contains(d.current)
}

func (d DFA) Rejecting() bool {
	return d.start == d.reject
}

func (d DFA) ReadCharacter(char string) {
	x := d.rules.GetFromState(char)
	if len(x) < 1 {
		d.current = d.reject
	}
	d.current = x[0][1]
}

func (d DFA) ReadString(word string) {
	for _, char := range word {
		if d.Rejecting() {
			return
		}
		d.ReadCharacter(string(char))
	}
}

func (d DFA) GetAllStates() Set {
	return d.states
}

func (d DFA) GetAllTransitionsFor(from, with string) (set Set) {
	if tran, ok := d.rules[from]; ok != false {
		if to, okk := tran[with]; okk != false {
			set.Add(to)
			return
		}
	}
	return
}

func (d DFA) GetAllEnds() (set Set) {
	for _, b := range d.rules {
		for _, j := range b {
			set.Add(j)
		}
	}
	return
}

func (d DFA) GetStartStates() Set {
	start := NewSet()
	start.Add(d.start)
	return start
}

func (d DFA) GetAcceptStates() Set {
	return d.accept
}

func (d DFA) GetReject() string {
	return d.reject
}

func (d DFA) GetAllRules() []Rule {
	return d.rules.GetAllRules()
}

func (d DFA) GetFromState(from string) [][2]string {
	return d.rules.GetFromState(from)
}
