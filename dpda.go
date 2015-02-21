package Automation

// Deterministic Pushdown Automata
type DPDA struct {
	start   string
	reject  string
	current string
	states  Set
	accept  Set
	stack   *Stack
	rules   DPRuleBook
}

// Returns new Deterministic Pushdown Automata.
func NewDPDA(start, reject, current string, states, accept Set, stack *Stack, rules DPRuleBook) *DPDA {
	return &DPDA{start: start, reject: reject, current: current, states: states, accept: accept, stack: stack, rules: rules}
}

func (d *DPDA) Restart() {
	d.current = d.start
	for i := 0; i < d.stack.Len()-1; i++ {
		_, _ = d.stack.Pop()
	}
}

func (d *DPDA) String() string {
	return "Current state: " + d.current + "\nState of the stack: " + d.stack.String()
}

// Whether the string so far is part of the language.
func (d *DPDA) Accepting() bool {
	return d.accept.Contains(d.current)
}

// Whether the string so far is not part of the language.
func (d *DPDA) Rejecting() bool {
	return d.current == d.reject
}

// Process of a single character at a time.
func (d *DPDA) ReadCharacter(char string) {
	stackTop, err := d.stack.Peek()
	pop := true

	if !err {
		d.current = d.reject
		return
	}

	rule := d.rules.GetRuleEnd(d.current, char, stackTop)
	if rule == nil {
		rule = d.rules.GetRuleEnd(d.current, char, stackTop+"!")
		if rule == nil {
			d.current = d.reject
			return
		}
		pop = false
	}

	d.current = rule[0][0]
	if pop == true {
		_, _ = d.stack.Pop()
	}

	if rule[0][1] != " " {
		d.stack.Push(rule[0][1])
	}

}

// Gets a string and process it a single character at a time.
func (d *DPDA) ReadString(word string) {
	for _, char := range word {
		if d.Rejecting() {
			return
		}
		d.ReadCharacter(string(char))
	}
}

// Returns the string representation of stack.
// The format is '[' member1 ',' member2 ',' ... ']'
func (d *DPDA) GetStackAsString() string {
	return d.stack.String()
}

// Returns the stack of the machine.
func (d *DPDA) GetStack() Stack {
	return *(d.stack)
}
