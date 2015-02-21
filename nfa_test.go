package Automation

import "testing"

func TestCreationNFA(t *testing.T) {
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

	for _, char := range "wna" {
		nfa.ReadCharacter(string(char))
	}

	if !nfa.Accepting() {
		t.Error("Must be accepting!")
	}

	nfa.ReadCharacter(string("9"))
	if !nfa.Rejecting() {
		t.Error("Must be rejectting!")
	}
	if nfa.Accepting() {
		t.Error("Must not be accepting!")
	}
}

func TestRecognitionNFA(t *testing.T) {
	rules := NewRules("from", "without", "toto")
	book := NewNRuleBook(rules)
	book.AddRules(NewRules("toto", "metal", "otot"))
	book.AddRules(NewRules("otot", "oxigen", "from"))
	book.AddRules(NewRules("from", "oxigen", "otot"))
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

	alpha := nfa.GetAlphabet().Values()
	if len(alpha) != 16 {
	}

	testSet := NewSet()
	for _, letter := range "metaloxtorwithoutnadowprionioxigengen" {
		testSet.Add(string(letter))
	}

	for _, letter := range alpha {
		if !testSet.Contains(letter) {
			t.Error("Problem with alphabet!")
			t.Error(letter + " was found")
		}
	}
}

func TestRestart(t *testing.T) {
	rules := NewRules("from", "without", "toto")
	book := NewNRuleBook(rules)
	book.AddRules(NewRules("toto", "metal", "otot"))
	book.AddRules(NewRules("otot", "oxigen", "from"))
	book.AddRules(NewRules("from", "oxigen", "otot"))
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

	nfa.ReadCharacter("w")
	nfa.Restart()
	if !nfa.current.Eq(nfa.start) {
		t.Error("Can't restart. Some problem.")
	}
}
