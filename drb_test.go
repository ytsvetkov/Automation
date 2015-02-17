package Automation

import "testing"

func TestAddInDeterministicRuleBook(t *testing.T) {
	rule1 := NewRule("a", "b", "c")
	rule2 := NewRule("b", "c", "a")
	rule3 := NewRule("c", "a", "b")

	book := NewEmptyDRuleBook()

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
		t.Error("Problem with third addition of rule !!!")
	}

	badrule := NewRule("a", "b", "w")
	err := book.AddRule(badrule)
	if err == nil {
		t.Error("Problem with detection of non-deterministic behaviour !!!")
	}
}

func TestDeterministicTransitions(t *testing.T) {
	rule1 := NewRule("a", "b", "c")
	rule2 := NewRule("b", "c", "a")
	rule3 := NewRule("c", "a", "b")

	book := NewEmptyDRuleBook()

	book.AddRule(rule1)
	book.AddRule(rule2)
	book.AddRule(rule3)

	state := book.GetFromState("a")
	if len(state) != 1 {
		t.Error("Problem with 'GetFromState': returns more than one transition!")
	}

	state = book.GetFromState("b")
	if len(state) != 1 {
		t.Error("Problem with 'GetFromState': returns more than one transition!")
	}

	state = book.GetFromState("c")
	if len(state) != 1 {
		t.Error("Problem with 'GetFromState': returns more than one transition!")
	}

	state = book.GetFromState("w")
	if len(state) != 0 {
		t.Error("Problem with 'GetFromState': returns transitions for non-existing states!")
	}
}

func TestDeterministicEndOfTransitions(t *testing.T) {
	rule1 := NewRule("a", "b", "c")
	rule4 := NewRule("a", "d", "w")
	rule2 := NewRule("b", "c", "a")
	rule3 := NewRule("c", "a", "b")

	book := NewEmptyDRuleBook()

	book.AddRule(rule1)
	book.AddRule(rule2)
	book.AddRule(rule3)
	book.AddRule(rule4)

	states := book.GetFromTransition("a", "d")
	if states.Cardinality() != 1 {
		t.Error("Incorect number of states!")
	}

	sliceStates := states.Values()
	for _, state := range sliceStates {
		if state != "w" {
			t.Error("Incorect states!")
		}
	}

}
