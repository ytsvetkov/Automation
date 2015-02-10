package Automation

import "testing"

func TestAddInNonDeterministicRuleBook(t *testing.T) {
	rule1 := NewRule("a", "b", "c")
	rule2 := NewRule("b", "c", "a")
	rule3 := NewRule("c", "a", "b")
	rule4 := NewRule("a", "b", "q")
	rule5 := NewRule("b", "c", "w")
	rule6 := NewRule("c", "a", "e")

	book := NewEmptyNRuleBook()

	err1 := book.AddRule(rule1)
	if err1 != nil {
		t.Error("Problem with initial addition of rule !!!")
	}
	err2 := book.AddRule(rule2)
	if err2 != nil {
		t.Error("Problem with second addition of rule !!!")
	}
	err3 := book.AddRule(rule3)
	if err3 != nil {
		t.Error("Problem with thirdth addition of rule !!!")
	}
	err4 := book.AddRule(rule4)
	if err4 != nil {
		t.Error("Problem with initial addition of rule !!!")
	}
	err5 := book.AddRule(rule5)
	if err5 != nil {
		t.Error("Problem with second addition of rule !!!")
	}
	err6 := book.AddRule(rule6)
	if err6 != nil {
		t.Error("Problem with thirdth addition of rule !!!")
	}
}

func TestNonDeterministicTransitions(t *testing.T) {
	rule1 := NewRule("a", "b", "c")
	rule2 := NewRule("b", "c", "a")
	rule3 := NewRule("c", "a", "b")
	rule4 := NewRule("a", "b", "q")
	rule5 := NewRule("b", "c", "w")
	rule6 := NewRule("c", "a", "e")

	slice := []*Rule{rule1, rule2, rule3, rule4, rule5, rule6}
	book := NewNRuleBook(slice)

	rule7 := NewRule("c", "a", "z")
	book.AddRule(rule7)

	states := book.GetFromState("c")
	if len(states) != 3 {
		t.Error("Problem with getting transitions!")
	}
	states = book.GetFromState("p")
	if len(states) != 0 {
		t.Error("Problem with getting transitions!")
	}
	states = book.GetFromState("b")
	if len(states) != 2 {
		t.Error("Problem with getting transitions!")
	}
}

func TestNonDeterministicEndOfTransitions(t *testing.T) {
	rule1 := NewRule("a", "b", "c")
	rule2 := NewRule("b", "c", "a")
	rule3 := NewRule("c", "a", "b")
	rule4 := NewRule("a", "b", "q")
	rule5 := NewRule("b", "c", "w")
	rule6 := NewRule("c", "a", "e")

	slice := []*Rule{rule1, rule2, rule3, rule4, rule5, rule6}
	book := NewNRuleBook(slice)

	rule7 := NewRule("c", "a", "z")
	book.AddRule(rule7)

	states := book.GetFromTransition("c", "a")
	if states.Cardinality() != 3 {
		t.Error("Incorect number of states!")
	}

	sliceStates := states.Values()
	for _, state := range sliceStates {
		if (state != "b") && (state != "e") && (state != "z") {
			t.Error("Incorect states!")
		}
	}
}
