package Automation

import "testing"

func ExamDFA() *DFA {
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

	book, _ := NewDRuleBook(rules)

	dfa := NewDFA("a", "c", states, accept, book)

	return dfa
}

func ExamNFA() *NFA {
	rules := NewRules("from", "without", "toto")
	book := NewNRuleBook(rules)
	book.AddRules(NewRules("toto", "metal", "otot"))
	book.AddRules(NewRules("otot", "oxigen", "from"))
	book.AddRules(NewRules("from", "mineral", "otot"))
	book.AddRules(NewRules("otot", "prion", "toto"))
	book.AddRules(NewRules("toto", "tornado", "from"))
	book.AddRule(NewRule("from", "w", "otot"))

	start := NewSet()
	start.Add("from")

	stt := []string{"from", "toto", "otot"}
	states := SetFromSlice(stt)

	acc := []string{"otot"}
	accept := SetFromSlice(acc)

	nfa := NewNFA(start, states, accept, "err", book)

	return nfa
}

func TestUnion(t *testing.T) {
	x := ExamNFA()
	y := ExamDFA()
	z := Union(x, y)

	z.ReadCharacter("w")
	if !z.Accepting() {
		t.Error("Should have accepted that!")
	}
	z.ReadCharacter("l")
	if !z.Accepting() {
		t.Error("Should have accepted that!")
	}
	z.ReadCharacter("5")
	if z.Accepting() {
		t.Error("Should not have accepted that!")
	}
	if !z.Rejecting() {
		t.Error("Should have rejected that!")
		t.Error(z.current.Contains(z.reject))
	}

	z.Restart()

	z.ReadString("1")
	if !z.Accepting() {
		t.Error("Should have accepted that!")
	}
}

func TestConcat(t *testing.T) {
	x := ExamNFA()
	y := ExamDFA()
	z := Concatenation(x, y)

	z.ReadString("wl11")

	if !z.Accepting() {
		t.Error("Should have accepted that!")
	}

	z.ReadCharacter("6")

	if z.Accepting() {
		t.Error("Should not have accepted that!")
	}
	if !z.Rejecting() {
		t.Error("Should have rejected that!")
	}
}
