package Automation

type Tape struct {
	left   []string
	middle string
	right  []string
}

func NewTape() *Tape {
	return &Tape{left: make([]string, 0), middle: "", right: make([]string, 0)}
}

func (t *Tape) MoveRight() {
	t.left = append(t.left, t.middle)
}

func (t *Tape) MoveLeft() {

}

func (t *Tape) Read() string {
	return t.middle
}

func (t *Tape) Write(str string) {
	t.middle = str
}
