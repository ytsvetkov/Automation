package Automation

type DPDA struct {
	start   string
	reject  string
	current string
	states  Set
	accept  Set
	stack   Stack
	rules   DPRuleBook
}

func NewDPDA(start, reject, current string, states, accept Set, stack Stack, rules DPRuleBook) DPDA {
	return DPDA{start: start, reject: reject, current: current, states: states, accept: accept, stack: stack, rules: rules}
}

func (d DPDA) Accepting() bool {
	return d.accept.Contains(d.current)
}

func (d DPDA) Rejecting() bool {
	return d.current == d.reject
}

func (d DPDA) ReadCharacter(char string) {
	stackTop, err := d.stack.Peek()
	if !err {
		d.start = d.reject
		return
	}
	x := d.rules.GetRuleEnd(d.current, char, stackTop)
	if x == nil {
		d.start = d.reject
		return
	}
	d.start = x[0][0]
	_, _ = d.stack.Pop()
	d.stack.Push(x[0][1])
}

func (d DPDA) ReadString(word string) {
	for _, char := range word {
		if d.Rejecting() {
			return
		}
		d.ReadCharacter(string(char))
	}
}
