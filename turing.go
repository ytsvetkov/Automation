package Automation

type TuringMachine struct {
	tape    Tape
	states  Set
	accept  Set
	reject  Set
	current string
	head    string //self.tape.middle
	stack   Stack
	rules   TuringRoolBook
}

func (t TuringMachine) Step() {

}

func (t TuringMachine) Run() {

}

func (t TuringMachine) String() string {

}
