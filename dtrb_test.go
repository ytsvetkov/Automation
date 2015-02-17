package Automation

import "testing"

func TestTuringRulebook(t *testing.T) {
	rule1 := NewTRule("alpha", "0", "beta", "147", "LEFT")
	rule2 := NewTRule("alpha", "1", "gama", "142", "RIGHT")
	rule3 := NewTRule("alpha", "2", "teta", "144", "NOP")

	book := NewEmptyTuringRoolBook()

	err1 := book.AddRule(rule1)
	if err1 != nil {
		t.Error("Problem with first addition!")
	}
	err1 = book.AddRule(rule2)
	if err1 != nil {
		t.Error("Problem with addition!")
	}
	err1 = book.AddRule(rule3)
	if err1 != nil {
		t.Error("Problem with addition!")
	}
	if len(book.GetAllRules()) != 3 {
		t.Error("Problem with initialisation!")
	}

	rule4 := NewTRule("alpha", "2", "teta", "144", "NOP")
	err := book.AddRule(rule4)
	if err == nil {
		t.Error("NonDeterminism not caught!")
	}

	if len(book.GetAllRules()) != 3 {
		t.Error("BigProblem")
	}
}
