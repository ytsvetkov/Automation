package Automation

import "testing"

func TestCreatingDTuring(t *testing.T) {
	rule1 := NewTRule("alpha", "0", "beta", "2", "LEFT")
	rule2 := NewTRule("alpha", "1", "gama", "0", "RIGHT")
	rule3 := NewTRule("alpha", "2", "teta", "1", "NOP")
	rule4 := NewTRule("beta", "1", "alpha", "2", "NOP")

	book := NewEmptyTuringRoolBook()
	book.AddRule(rule1)
	book.AddRule(rule2)
	book.AddRule(rule3)
	book.AddRule(rule4)

	accept := NewSet()
	accept.Add("gama")

	states := NewSet()
	states.Add("alpha")
	states.Add("beta")
	states.Add("gama")
	states.Add("teta")

	tape := NewNonEmptyTape("0102&", "0", "0121012")

	tur := NewDTuringMachine("alpha", "teta", states, accept, tape, book)

	tur.Step()
	if tur.current != "beta" {
		t.Error("In the wrong state!")
	}

	if tur.GetTapeString() != "0102&20121012" {
		t.Error("Wrong state of the tape!")
	}

	tur.Step()
	if tur.Accepting() {
		t.Error("Should not be accepting. Should be rejecting!")
	}
	if !tur.Rejecting() {
		t.Error("Rejectiong recognition failed!")
	}
}

func TestRunTuringRun(t *testing.T) {
	rule1 := NewTRule("alpha", "0", "alpha", "1", "RIGHT")
	rule2 := NewTRule("alpha", "1", "alpha", "0", "RIGHT")
	rule3 := NewTRule("alpha", "#", "sigma", " ", "NOP")

	book := NewEmptyTuringRoolBook()
	book.AddRule(rule1)
	book.AddRule(rule2)
	book.AddRule(rule3)

	accept := NewSet()
	accept.Add("sigma")

	states := NewSet()
	states.Add("alpha")
	states.Add("sigma")
	states.Add("teta")

	tape := NewNonEmptyTape("", "0", "01010001001101110#")

	tur := NewDTuringMachine("alpha", "teta", states, accept, tape, book)

	tur.Run()
	if tur.Rejecting() {
		t.Error("Should not be rejecting!")
	}
	if !tur.Accepting() {
		t.Error("Should be accepting!")
	}
}
