package Automation

import "testing"

func TestDFACreationAndWork(t *testing.T) {
	rules := make([]*Rule, 0)
	rules = append(rules, NewRule("a", "0", "a"))
	rules = append(rules, NewRule("a", "1", "b"))
	rules = append(rules, NewRule("b", "0", "b"))
	rules = append(rules, NewRule("b", "1", "a"))
	rules = append(rules, NewRule("b", "2", "c"))
	rules = append(rules, NewRule("a", "2", "c"))
	rules = append(rules, NewRules("c", "012", "c")...)

	states := NewSet()
	states.Add("a")
	states.Add("b")
	states.Add("c")

	acc := []string{"a", "b"}
	accept := SetFromSlice(acc)

	rulebook, err := NewDRuleBook(rules)
	if err != nil {
		t.Error("Some problem with the Rulebook.")
	}

	dfa := NewDFA("a", "c", states, accept, rulebook)
	if dfa.current != "a" {
		t.Error("Problem with starting state")
	}
	if dfa.start != "a" {
		t.Error("Problem with starting state")
	}

	for _, char := range "10101010101101010" {
		dfa.ReadCharacter(string(char))
	}

	if dfa.current != "b" {
		t.Error("Problem with transitioning!")
	}

	if dfa.Rejecting() {
		t.Error("Problem with rejection detection!")
	}

	dfa.ReadCharacter("6")
	if !dfa.Rejecting() {
		t.Error("Problem with rejection detection!")
	}

	dfa.Restart()
	if dfa.current != "a" {
		t.Error("Problem with starting state")
	}
	if dfa.start != "a" {
		t.Error("Problem with starting state")
	}

	dfa.ReadCharacter("2")
	if !dfa.Rejecting() {
		t.Error("Problem with rejection detection!")
	}
}

func TestAlphabetDFA(t *testing.T) {
	rules := make([]*Rule, 0)
	rules = append(rules, NewRule("a", "0", "a"))
	rules = append(rules, NewRule("a", "1", "b"))
	rules = append(rules, NewRule("b", "0", "b"))
	rules = append(rules, NewRule("b", "1", "a"))
	rules = append(rules, NewRule("b", "2", "c"))
	rules = append(rules, NewRule("a", "2", "c"))
	rules = append(rules, NewRules("c", "012", "c")...)

	states := NewSet()
	states.Add("a")
	states.Add("b")
	states.Add("c")

	acc := []string{"a", "b"}
	accept := SetFromSlice(acc)

	rulebook, _ := NewDRuleBook(rules)

	dfa := NewDFA("a", "c", states, accept, rulebook)

	alpha := dfa.GetAlphabet().Values()
	if len(alpha) != 3 {
		t.Error("Problem with alphabet!")
	}
	for _, letter := range alpha {
		if letter != "0" && letter != "1" && letter != "2" {
			t.Error("Problem with alphabet!")
		}
	}

	dfa.ReadString("10101010101101010")
	if dfa.current != "b" {
		t.Error("Problem with transitioning!")
	}

	if dfa.Rejecting() {
		t.Error("Problem with rejection detection!")
	}
	if !dfa.Accepting() {
		t.Error("Problem with rejection detection!")
	}

}
