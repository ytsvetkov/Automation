package Automation

import "testing"

func TestTuringNRulebook(t *testing.T) {
	rule1 := NewTRule("alpha", "0", "beta", "147", "LEFT")
	rule2 := NewTRule("alpha", "0", "gama", "142", "RIGHT")
	rule3 := NewTRule("alpha", "0", "delta", "144", "NOP")
	rule4 := NewTRule("alpha", "0", "psi", "144", "NOP")
	rule5 := NewTRule("alpha", "0", "ksi", "144", "NOP")
	rule6 := NewTRule("alpha", "0", "zeta", "144", "NOP")

	book := NewEmptyNTuringNRoolBook()
	book.AddRule(rule1)
	book.AddRule(rule2)
	book.AddRule(rule3)
	book.AddRule(rule4)
	book.AddRule(rule5)
	book.AddRule(rule6)

	if len(book.GetAllRules()) != 6 {
		t.Error("Problem with initialisation!")
	}
}

func TestNonDeterministicTuringTransitions(t *testing.T) {
	rule1 := NewTRule("alpha", "0", "beta", "147", "LEFT")
	rule2 := NewTRule("alpha", "0", "gama", "142", "RIGHT")
	rule3 := NewTRule("alpha", "0", "delta", "144", "NOP")
	rule4 := NewTRule("alpha", "0", "psi", "144", "NOP")
	rule5 := NewTRule("alpha", "0", "ksi", "144", "NOP")
	rule6 := NewTRule("alpha", "0", "zeta", "144", "NOP")

	book := NewEmptyNTuringNRoolBook()
	book.AddRule(rule1)
	book.AddRule(rule2)
	book.AddRule(rule3)
	book.AddRule(rule4)
	book.AddRule(rule5)
	book.AddRule(rule6)

	state := book.GetFromState("alpha")
	if len(state) != 6 {
		t.Error("Problem with getting transitions!")
	}

	rule7 := NewTRule("alpha", "0", "tete", "144", "NOP")
	rule8 := NewTRule("alpha", "0", "eta", "144", "NOP")
	book.AddRule(rule7)
	book.AddRule(rule8)

	state = book.GetFromState("alpha")
	if len(state) != 8 {
		t.Error("Problem with getting transitions!")
	}
}

func TestNonDeterministicTuringRuleEnd(t *testing.T) {
	rule1 := NewTRule("alpha", "0", "beta", "147", "LEFT")
	rule2 := NewTRule("alpha", "0", "gama", "142", "RIGHT")
	rule3 := NewTRule("alpha", "0", "delta", "144", "NOP")
	rule4 := NewTRule("alpha", "1", "psi", "144", "NOP")
	rule5 := NewTRule("alpha", "1", "ksi", "144", "NOP")
	rule6 := NewTRule("alpha", "1", "zeta", "144", "NOP")

	book := NewEmptyNTuringNRoolBook()
	book.AddRule(rule1)
	book.AddRule(rule2)
	book.AddRule(rule3)
	book.AddRule(rule4)
	book.AddRule(rule5)
	book.AddRule(rule6)

	tran := book.GetRuleEnd("alpha", "1")
	if len(tran) != 3 {
		t.Error("Problem with 'GetRuleEnd': returns more than one transition!")
	}

	tran = book.GetRuleEnd("alpha", "0")
	if len(tran) != 3 {
		t.Error("Problem with 'GetRuleEnd': returns more than one transition!")
	}

	rule7 := NewTRule("alpha", "0", "tete", "144", "NOP")
	rule8 := NewTRule("alpha", "0", "eta", "144", "NOP")
	book.AddRule(rule7)
	book.AddRule(rule8)

	tran = book.GetRuleEnd("alpha", "0")
	if len(tran) != 5 {
		t.Error("Problem with 'GetRuleEnd': returns more than one transition!")
	}

}

func TestNonDeterministicTuringFromTransition(t *testing.T) {
	rule1 := NewTRule("alpha", "0", "beta", "147", "LEFT")
	rule2 := NewTRule("alpha", "0", "gama", "142", "RIGHT")
	rule3 := NewTRule("alpha", "0", "delta", "144", "NOP")
	rule4 := NewTRule("alpha", "1", "psi", "144", "NOP")
	rule5 := NewTRule("alpha", "1", "ksi", "144", "NOP")
	rule6 := NewTRule("alpha", "1", "zeta", "144", "NOP")

	book := NewEmptyNTuringNRoolBook()
	book.AddRule(rule1)
	book.AddRule(rule2)
	book.AddRule(rule3)
	book.AddRule(rule4)
	book.AddRule(rule5)
	book.AddRule(rule6)

	tran := book.GetFromTransition("alpha")
	if tran.Cardinality() != 6 {
		t.Error("Problem with 'GetFromTransition': returns different number of states!")
	}

	rule7 := NewTRule("alpha", "0", "tete", "144", "NOP")
	rule8 := NewTRule("alpha", "0", "eta", "144", "NOP")
	book.AddRule(rule7)
	book.AddRule(rule8)

	tran = book.GetFromTransition("alpha")
	if tran.Cardinality() != 8 {
		t.Error("Problem with 'GetFromTransition': returns different number of states!")
	}
}

func TestNonDeterministicTuringGetAllRules(t *testing.T) {
	rule1 := NewTRule("alpha", "0", "beta", "147", "LEFT")
	rule2 := NewTRule("alpha", "0", "gama", "142", "RIGHT")
	rule3 := NewTRule("alpha", "0", "delta", "144", "NOP")
	rule4 := NewTRule("alpha", "1", "psi", "144", "NOP")
	rule5 := NewTRule("alpha", "1", "ksi", "144", "NOP")
	rule6 := NewTRule("alpha", "1", "zeta", "144", "NOP")

	book := NewEmptyNTuringNRoolBook()
	book.AddRule(rule1)
	book.AddRule(rule2)
	book.AddRule(rule3)
	book.AddRule(rule4)
	book.AddRule(rule5)
	book.AddRule(rule6)

	tran := book.GetAllRules()
	if len(tran) != 6 {
		t.Error("Problem with 'GetAllRules'!")
	}
}
