package Automation

type NTuringMachine struct {
	tape      []Set
	states    Set
	accept    Set
	reject    string
	current   Set
	rules     TuringNRoolBook
	headIndex []int
}

func NewNTuringMachine(current, states, accept Set, reject string, tape *Tape, rules TuringNRoolBook) *NTuringMachine {
	s := NewSet()
	t := []Set{}
	for _, letter := range tape.left {
		s.Add(string(letter))
		t = append(t, s)
		s = NewSet()
	}
	s.Add(string(tape.middle))
	t = append(t, s)
	s = NewSet()
	for _, letter := range tape.right {
		s.Add(string(letter))
		t = append(t, s)
		s = NewSet()
	}
	return &NTuringMachine{current: current, states: states, accept: accept, reject: reject,
		tape: t, rules: rules, headIndex: []int{len(tape.left)}}
}

func (t *NTuringMachine) Accepting() bool {
	return (!t.Rejecting()) && (t.accept.Intersection(t.current).Cardinality() != 0)
}

func (t *NTuringMachine) Rejecting() bool {
	return t.current.Contains(t.reject)
}

func (t *NTuringMachine) Step() {
	transition := [][3]string{}

	for _, index := range t.headIndex {
		for _, letter := range t.tape[index].Values() {
			for _, state := range t.current.Values() {
				transition = append(transition, t.rules.GetRuleEnd(state, letter)...)
			}
		}
	}

	if len(transition) < 1 {
		t.current.Add(t.reject)
		return
	}

	t.current = NewSet()
	for _, headIndex := range t.headIndex {
		t.tape[headIndex] = NewSet()
		for _, tran := range transition {
			t.current.Add(tran[0])
			t.tape[headIndex].Add(tran[1])
		}
	}

	newHead := []int{}
	for _, headIndex := range t.headIndex {
		for _, tran := range transition {
			switch tran[2] {
			case "LEFT":
				newHead = append(newHead, headIndex-1)
			case "RIGHT":
				newHead = append(newHead, headIndex+1)
			}
		}
	}
	t.headIndex = newHead
}

func (t *NTuringMachine) Run() {
	for !t.Rejecting() && !t.Accepting() {
		t.Step()
	}
}

func (t *NTuringMachine) GetCurrentState() string {
	return t.current.String()
}

func (t *NTuringMachine) String() string {
	str := "Tape: " + t.GetTapeString() + "\n"
	str += "Head: " + t.tape[t.headIndex[0]].String() + "\n"
	str += "All States: " + t.states.String() + "\n"
	str += "Current State: " + t.current.String() + "\n"
	str += "Accept States: " + t.accept.String() + "\n"
	str += "Reject States: " + t.reject + "\n"
	str += "Rules: " + t.rules.String() + "\n"
	return str
}

func (t *NTuringMachine) GetTapeString() string {
	head := false
	str := "["

	for i := 0; i < len(t.tape)-1; i++ {
		for _, j := range t.headIndex {
			if i == j {
				str += "(" + t.tape[i].String() + "), "
				head = true
			}
		}
		if !head {
			str += t.tape[i].String() + ", "
		}
		head = false
	}
	if len(t.tape) != 0 {
		return str + t.tape[len(t.tape)-1].String() + "]"
	}
	return str + "]"
}
