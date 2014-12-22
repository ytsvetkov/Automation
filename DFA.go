package Automation

type DFA struct {
	start   string
	current string
	accept  *Set
	reject  string
	rules   DeterministicRuleBook
}

func NewDFA(start, reject string, accept []string, rules DeterministicRuleBook) *DFA {
	accepting := NewSet()
	accepting.AddSlice(accept)
	return &DFA{start: start, current: start, accept: accepting, reject: reject, rules: rules}
}

func (d *DFA) Accepting() bool {
	return d.accept.Contains(d.current)
}

func (d *DFA) Rejecting() bool {
	return d.reject == d.current
}

func (d *DFA) ReadCharacter(char string) {
	rule := d.rules.GetRule(d.current, char)
	if rule == nil {
		d.current = d.reject
	} else {
		d.current = rule.GetEndingState()
	}
}

func (d *DFA) ReadString(word string) {
	for i := 0; !d.Rejecting() && i < len(word); i++ {
		d.ReadCharacter(string(word[i]))
	}
}
