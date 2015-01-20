package Automation

type NPDA struct {
	start   Set
	reject  Set
	current Set
	states  Set
	accept  Set
	stack   Stack
	rules   DPRuleBook
}

// func NewNPDA(start, reject, current string, states, accept Set, stack Stack, rules DPRuleBook) NPDA {
// 	return NPDA{start: start, reject: reject, current: current, states: states, accept: accept, stack: stack, rules: rules}
// }

// func (d NPDA) Accepting() bool {
// 	return d.accept.Contains(d.current)
// }

// func (d NPDA) Rejecting() bool {
// 	return d.current == d.reject
// }

// func (d NPDA) ReadCharacter(char string) {
// 	stackTop, err := d.stack.Peek()
// 	if !err {
// 		d.start = d.reject
// 		return
// 	}
// 	x := d.rules.GetRuleEnd(d.current, char, stackTop)
// 	if x == nil {
// 		d.start = d.reject
// 		return
// 	}
// 	d.start = x[0][0]
// 	_, _ = d.stack.Pop()
// 	d.stack.Push(x[0][1])
// }

// func (d NPDA) ReadString(word string) {
// 	for _, char := range word {
// 		if d.Rejecting() {
// 			return
// 		}
// 		d.ReadCharacter(string(char))
// 	}
// }
