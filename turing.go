package Automation

// Turing machine
type TuringMachine struct {
	tape    *Tape
	states  Set
	accept  Set
	reject  string
	current string
	rules   TuringRoolBook
}

func NewDTuringMachine(current, reject string, states, accept Set, tape *Tape, rules TuringRoolBook) *TuringMachine {
	return &TuringMachine{current: current, states: states, accept: accept, reject: reject, tape: tape, rules: rules}
}

// Returns whether the string so far is part of the language.
func (t *TuringMachine) Accepting() bool {
	return t.accept.Contains(t.current)
}

// Returns whether the string so far is not part of the language.
func (t *TuringMachine) Rejecting() bool {
	return t.current == t.reject
}

// One step transitioning.
func (t *TuringMachine) Step() {
	transition := t.rules.GetRuleEnd(t.current, t.tape.Read())
	if len(transition) < 1 {
		t.current = t.reject
		return
	}

	t.current = transition[0][0]
	t.tape.Write(transition[0][1])

	switch transition[0][2] {
	case "LEFT":
		t.tape.MoveLeft()
	case "RIGHT":
		t.tape.MoveRight()
	}
}

// Starts the machine until it halts.
func (t *TuringMachine) Run() {
	for !t.Rejecting() && !t.Accepting() {
		t.Step()
	}
}

// Returns the current state.
func (t *TuringMachine) GetCurrentState() string {
	return t.current
}

func (t *TuringMachine) String() string {
	str := "Tape: " + t.tape.String() + "\n"
	str += "All States: " + t.states.String() + "\n"
	str += "Current State: " + t.current + "\n"
	str += "Accept States: " + t.accept.String() + "\n"
	str += "Reject States: " + t.reject + "\n"
	str += "Rules: " + t.rules.String() + "\n"
	return str
}

// String representation of the tape of the machine.
func (t *TuringMachine) GetTapeString() string {
	return t.tape.left + t.tape.middle + t.tape.right
}
