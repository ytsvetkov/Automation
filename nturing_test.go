package Automation

import "testing"

func TestCreationNTuring(t *testing.T) {
	rule1 := NewTRule("alpha", "0", "beta", "7", "LEFT")
	rule2 := NewTRule("alpha", "1", "gama", "2", "RIGHT")
	rule3 := NewTRule("alpha", "0", "delta", "4", "NOP")
	rule4 := NewTRule("alpha", "1", "psi", "1", "NOP")
	rule5 := NewTRule("alpha", "1", "ksi", "4", "NOP")
	rule6 := NewTRule("alpha", "0", "zeta", "1", "RIGHT")

	book := NewEmptyNTuringNRoolBook()
	book.AddRule(rule1)
	book.AddRule(rule2)
	book.AddRule(rule3)
	book.AddRule(rule4)
	book.AddRule(rule5)
	book.AddRule(rule6)

	states := NewSet()
	states.Add("alpha")
	states.Add("beta")
	states.Add("gama")
	states.Add("delta")
	states.Add("ksi")
	states.Add("psi")
	states.Add("zeta")

	acc := NewSet()
	acc.Add("ksi")

	current := NewSet()
	current.Add("alpha")

	tape := NewNonEmptyTape("abcd", "0", "efgh")

	machine := NewNTuringMachine(current, acc, states, "err", tape, book)

	machine.Step()
	other := NewSet()
	other.Add("beta")
	other.Add("delta")
	other.Add("zeta")
	if machine.current.Intersection(other).Cardinality() != 3 {
		t.Error("Not the correct set of current states!")
	}

	machine.Step()
	if !machine.Rejecting() {
		t.Error("Should be rejecting!")
	}
	if machine.Accepting() {
		t.Error("Should not be accepting!")
	}

}
