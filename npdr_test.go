package Automation

import "testing"

func TestAddInNonDeterministicPushdownRuleBook(t *testing.T) {
	rule1 := NewPRule("a", "1", "#", "a", "0")
	rule2 := NewPRule("a", "1", "0", "a", "0")
	rule3 := NewPRule("a", "1", "1", "a", "0")
	rule4 := NewPRule("a", "1", "#", "b", "1")
	rule5 := NewPRule("a", "1", "0", "b", "1")
	rule6 := NewPRule("a", "1", "1", "b", "1")

	book := NewEmptyNPRuleBook()
	book.AddRule(rule1)
	book.AddRule(rule2)
	book.AddRule(rule3)
	book.AddRule(rule4)
	book.AddRule(rule5)
	book.AddRule(rule6)

	rule7 := NewPRule("b", "1", "1", "b", "1")
	rule8 := NewPRule("b", "1", "1", "b", "2")
	book.AddRule(rule7)
	book.AddRule(rule8)

	if len(book.GetAllRules()) != 8 {
		t.Error("Problem with addition of rules !!!")
	}
}

func TestNonDeterministicPushdownTransitions(t *testing.T) {
	rule1 := NewPRule("a", "1", "#", "a", "0")
	rule2 := NewPRule("a", "1", "0", "a", "0")
	rule3 := NewPRule("a", "1", "1", "a", "0")
	rule4 := NewPRule("a", "1", "#", "b", "1")
	rule5 := NewPRule("a", "1", "0", "b", "1")
	rule6 := NewPRule("a", "1", "1", "b", "1")

	book := NewEmptyNPRuleBook()
	book.AddRule(rule1)
	book.AddRule(rule2)
	book.AddRule(rule3)
	book.AddRule(rule4)
	book.AddRule(rule5)
	book.AddRule(rule6)

	state := book.GetFromState("a")
	if len(state) != 6 {
		t.Error("Problem with 'GetFromState': returns different number of transitions!")
	}

	rule7 := NewPRule("b", "1", "0", "b", "1")
	rule8 := NewPRule("b", "1", "0", "b", "2")
	book.AddRule(rule7)
	book.AddRule(rule8)

	state = book.GetFromState("b")
	if len(state) != 2 {
		t.Error("Problem with 'GetFromState': returns different number of transitions!")
	}

	state = book.GetFromState("w")
	if len(state) != 0 {
		t.Error("Problem with 'GetFromState': returns transitions for non-existing states!")
	}
}

func TestNonDeterministicPushdownRuleEnd(t *testing.T) {
	rule1 := NewPRule("a", "1", "#", "a", "0")
	rule2 := NewPRule("a", "1", "0", "a", "0")
	rule3 := NewPRule("a", "1", "1", "a", "0")
	rule4 := NewPRule("a", "1", "#", "b", "1")
	rule5 := NewPRule("a", "1", "0", "b", "1")
	rule6 := NewPRule("a", "1", "1", "b", "1")

	book := NewEmptyNPRuleBook()
	book.AddRule(rule1)
	book.AddRule(rule2)
	book.AddRule(rule3)
	book.AddRule(rule4)
	book.AddRule(rule5)
	book.AddRule(rule6)

	tran := book.GetRuleEnd("a", "1", "#")
	if len(tran) != 2 {
		t.Error("Problem with 'GetRuleEnd': returns more than one transition!")
	}

	rule7 := NewPRule("b", "1", "0", "b", "1")
	rule8 := NewPRule("b", "1", "0", "b", "2")
	book.AddRule(rule7)
	book.AddRule(rule8)

	tran = book.GetRuleEnd("b", "1", "0")
	if len(tran) != 2 {
		t.Error("Problem with 'GetRuleEnd': returns more than one transition!")
	}

	tran = book.GetRuleEnd("b", "3", "1")
	if len(tran) != 0 {
		t.Error("Problem with 'GetRuleEnd': returns more than one transition!")
	}
}

func TestNonDeterministicPushdownFromTransition(t *testing.T) {
	rule1 := NewPRule("a", "1", "#", "a", "0")
	rule2 := NewPRule("a", "1", "0", "c", "0")
	rule3 := NewPRule("a", "1", "1", "d", "0")
	rule4 := NewPRule("a", "1", "#", "b", "1")
	rule5 := NewPRule("a", "1", "0", "b", "1")
	rule6 := NewPRule("a", "1", "1", "b", "1")
	rule7 := NewPRule("b", "1", "0", "b", "1")
	rule8 := NewPRule("b", "1", "0", "a", "2")

	book := NewEmptyNPRuleBook()
	book.AddRule(rule1)
	book.AddRule(rule2)
	book.AddRule(rule3)
	book.AddRule(rule4)
	book.AddRule(rule5)
	book.AddRule(rule6)
	book.AddRule(rule7)
	book.AddRule(rule8)

	tran := book.GetFromTransition("a")
	if tran.Cardinality() != 4 {
		t.Error("Problem with 'GetFromTransition': returns different number of states!")
	}

	tran = book.GetFromTransition("b")
	if tran.Cardinality() != 2 {
		t.Error("Problem with 'GetFromTransition': returns different number of states!")
	}

	tran = book.GetFromTransition("c")
	if tran.Cardinality() != 0 {
		t.Error("Problem with 'GetRuleEnd': returns transitions!")
	}
}

func TestNonDeterministicPushdownGetAllRules(t *testing.T) {
	rule1 := NewPRule("a", "1", "#", "a", "0")
	rule2 := NewPRule("a", "1", "0", "c", "0")
	rule3 := NewPRule("a", "1", "1", "d", "0")
	rule4 := NewPRule("a", "1", "#", "b", "1")
	rule5 := NewPRule("a", "1", "0", "b", "1")
	rule6 := NewPRule("a", "1", "1", "b", "1")
	rule7 := NewPRule("b", "1", "0", "b", "1")
	rule8 := NewPRule("b", "1", "0", "a", "2")

	book := NewEmptyNPRuleBook()
	book.AddRule(rule1)
	book.AddRule(rule2)
	book.AddRule(rule3)
	book.AddRule(rule4)
	book.AddRule(rule5)
	book.AddRule(rule6)
	book.AddRule(rule7)
	book.AddRule(rule8)

	tran := book.GetAllRules()
	if len(tran) != 8 {
		t.Error("Problem with 'GetAllRules'!")
	}
}
