package Automation

import "testing"

func TestAddInDeterministicPushdownRuleBook(t *testing.T) {
	rule1 := NewPRule("a", "1", "#", "a", "0")
	rule2 := NewPRule("a", "1", "0", "a", "0")
	rule3 := NewPRule("a", "1", "1", "a", "0")
	rule4 := NewPRule("a", "0", "#", "a", "1")
	rule5 := NewPRule("a", "0", "0", "a", "1")
	rule6 := NewPRule("a", "0", "1", "a", "1")

	book := NewEmptyDPRuleBook()

	err1 := book.AddRule(rule1)
	if err1 != nil {
		t.Error("Problem with initial addition of rule !!!")
	}
	err2 := book.AddRule(rule2)
	if err2 != nil {
		t.Error("Problem with initial addition of rule !!!")
	}
	err3 := book.AddRule(rule3)
	if err3 != nil {
		t.Error("Problem with initial addition of rule !!!")
	}
	err4 := book.AddRule(rule4)
	if err4 != nil {
		t.Error("Problem with initial addition of rule !!!")
	}
	err5 := book.AddRule(rule5)
	if err5 != nil {
		t.Error("Problem with initial addition of rule !!!")
	}
	err6 := book.AddRule(rule6)
	if err6 != nil {
		t.Error("Problem with initial addition of rule !!!")
	}

	badrule1 := NewPRule("a", "0", "1", "a", "2")
	badrule2 := NewPRule("a", "0", "1", "b", "2")

	err1 = book.AddRule(badrule1)
	if err1 == nil {
		t.Error("Problem with detection of non-deterministic behaviour !!!")
	}
	err2 = book.AddRule(badrule2)
	if err2 == nil {
		t.Error("Problem with detection of non-deterministic behaviour !!!")
	}
}

func TestDeterministicPushdownTransitions(t *testing.T) {
	rule1 := NewPRule("a", "1", "#", "a", "0")
	rule2 := NewPRule("a", "1", "0", "a", "0")
	rule3 := NewPRule("a", "1", "1", "a", "0")
	rule4 := NewPRule("a", "0", "#", "a", "1")
	rule5 := NewPRule("a", "0", "0", "a", "1")
	rule6 := NewPRule("a", "0", "1", "a", "1")

	book := NewEmptyDPRuleBook()
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

	rule7 := NewPRule("b", "0", "1", "a", "1")
	rule8 := NewPRule("b", "1", "1", "a", "1")
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

func TestDeterministicPushdownRuleEnd(t *testing.T) {
	rule1 := NewPRule("a", "1", "#", "a", "0")
	rule2 := NewPRule("a", "1", "0", "a", "0")
	rule3 := NewPRule("a", "1", "1", "a", "0")
	rule4 := NewPRule("a", "0", "#", "a", "1")
	rule5 := NewPRule("a", "0", "0", "a", "1")
	rule6 := NewPRule("a", "0", "1", "a", "1")

	book := NewEmptyDPRuleBook()
	book.AddRule(rule1)
	book.AddRule(rule2)
	book.AddRule(rule3)
	book.AddRule(rule4)
	book.AddRule(rule5)
	book.AddRule(rule6)

	tran := book.GetRuleEnd("a", "1", "#")
	if len(tran) != 1 {
		t.Error("Problem with 'GetRuleEnd': returns more than one transition!")
	}

	rule7 := NewPRule("b", "0", "1", "a", "1")
	rule8 := NewPRule("b", "1", "1", "a", "1")
	book.AddRule(rule7)
	book.AddRule(rule8)

	tran = book.GetRuleEnd("b", "1", "1")
	if len(tran) != 1 {
		t.Error("Problem with 'GetRuleEnd': returns more than one transition!")
	}

	tran = book.GetRuleEnd("b", "3", "1")
	if len(tran) != 0 {
		t.Error("Problem with 'GetRuleEnd': returns more than one transition!")
	}
}

func TestDeterministicPushdownFromTransition(t *testing.T) {
	rule1 := NewPRule("a", "1", "#", "a", "0")
	rule2 := NewPRule("a", "1", "0", "a", "0")
	rule3 := NewPRule("a", "1", "1", "a", "0")
	rule4 := NewPRule("a", "0", "#", "a", "1")
	rule5 := NewPRule("a", "0", "0", "a", "1")
	rule6 := NewPRule("a", "0", "1", "a", "1")
	rule7 := NewPRule("b", "0", "1", "b", "1")
	rule8 := NewPRule("b", "1", "1", "a", "1")

	book := NewEmptyDPRuleBook()
	book.AddRule(rule1)
	book.AddRule(rule2)
	book.AddRule(rule3)
	book.AddRule(rule4)
	book.AddRule(rule5)
	book.AddRule(rule6)
	book.AddRule(rule7)
	book.AddRule(rule8)

	tran := book.GetFromTransition("a")
	if tran.Cardinality() != 1 {
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

func TestDeterministicPushdownGetAllRules(t *testing.T) {
	rule1 := NewPRule("a", "1", "#", "a", "0")
	rule2 := NewPRule("a", "1", "0", "a", "0")
	rule3 := NewPRule("a", "1", "1", "a", "0")
	rule4 := NewPRule("a", "0", "#", "a", "1")
	rule5 := NewPRule("a", "0", "0", "a", "1")
	rule6 := NewPRule("a", "0", "1", "a", "1")
	rule7 := NewPRule("b", "0", "1", "b", "1")
	rule8 := NewPRule("b", "1", "1", "a", "1")

	book := NewEmptyDPRuleBook()
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
