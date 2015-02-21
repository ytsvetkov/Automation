package Automation

// List-like structure for the turing machine.
type Tape struct {
	left   string
	middle string
	right  string
}

func NewTape() *Tape {
	return &Tape{left: "", middle: " ", right: ""}
}

func NewNonEmptyTape(left string, middle string, right string) *Tape {
	return &Tape{left: left, middle: middle, right: right}
}

// Moves the head one cell to the right.
func (t *Tape) MoveRight() {
	t.left = t.left + t.middle

	if len(t.right) > 0 {
		t.middle = string((t.right)[0])
	} else {
		t.middle = " "
	}

	if len(t.right) > 0 {
		t.right = t.right[1:]
	} else {
		t.right = ""
	}

}

// Moves the head one cell to the left.
func (t *Tape) MoveLeft() {
	t.right = t.middle + t.right

	if len(t.left) != 0 {
		t.middle = string(t.left[len(t.left)-1])
		t.left = t.left[:len(t.left)-1]
	} else {
		t.middle = " "
		t.left = ""
	}
}

// Returns the symbol, pointer by the head.
func (t *Tape) Read() string {
	return t.middle
}

// Writes a symbol in the cell, refered by the head.
func (t *Tape) Write(str string) {
	t.middle = str
}

func (t *Tape) String() string {
	return "<" + t.left + ">(" + t.middle + ")<" + t.right + ">"
}
