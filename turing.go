package Automation

type TuringMachine struct {
	tape    *Tape
	states  Set
	accept  Set
	reject  Set
	current string
	rules   TuringRoolBook
}

func (t TuringMachine) Accepting() bool {
	return t.accept.Contains(t.current)
}

func (t TuringMachine) Rejecting() bool {
	return t.reject.Contains(t.current)
}

func (t TuringMachine) Step() {
	transition := t.rules.GetRuleEnd(t.current, t.tape.Read())
	t.current = transition[0][0]
	t.tape.Write(transition[0][1])

	switch transition[0][2] {
	case "Left":
		t.tape.MoveLeft()
	case "Right":
		t.tape.MoveRight()
	}
}

func (t TuringMachine) Run() {
	for !t.Accepting() && !t.Rejecting() {
		t.Step()
	}
}

func (t TuringMachine) String() string {
	str := "Tape: " + t.tape.String() + "\n"
	str += "Current State: " + t.current + "\n"
	str += "Accept States: " + t.accept.String() + "\n"
	str += "Reject States: " + t.reject.String() + "\n"
	str += "Rules: " + t.rules.String() + "\n"
	return str
}
