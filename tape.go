package Automation

type Tape struct {
	left   []string
	middle string
	right  []string
}

func NewTape() *Tape {
	return &Tape{left: make([]string, 0), middle: " ", right: make([]string, 0)}
}

func NewNonEmptyTape(left []string, middle string, right []string) *Tape {
	return &Tape{left: left, middle: middle, right: right}
}

func (t *Tape) MoveRight() {
	t.left = append(t.left, t.middle)

	if len(t.right) > 0 {
		t.middle = t.right[0]
		t.right = t.right[1:]
	} else {
		t.middle = " "
	}

}

func (t *Tape) MoveLeft() {
	temp := []string{t.middle}
	t.right = append(temp, t.right...)

	if len(t.left) > 0 {
		t.middle = t.left[len(t.left)-1]
		t.left = t.left[0:]
	} else {
		t.middle = " "
	}

}

func (t *Tape) Read() string {
	return t.middle
}

func (t *Tape) Write(str string) {
	t.middle = str
}

func (t *Tape) String() string {
	str := ""
	for _, letter := range t.left {
		str += string(letter)
	}
	str = str + "(" + t.middle + ")"
	for _, letter := range t.right {
		str += string(letter)
	}
	return str
}
