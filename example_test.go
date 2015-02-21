package Automation

func ExampleRuleAndRulebook() {
	a_to_b_with_q := NewRule("a", "q", "b")
	c_to_d_with_theAlphabet := NewRules("c", "abcdefghijklmnopqrstuvwxyz", "d")

	empty_rulebook := NewEmptyDRuleBook()
	empty_rulebook.AddRule(a_to_b_with_q)
	empty_rulebook.AddRules(c_to_d_with_theAlphabet)

	nonempty_rulebook := NewDRuleBook(c_to_d_with_theAlphabet)
	fmt.Println(nonempty_rulebook)
}

func ExampleDFA() *DFA {
	rules := make([]Rule, 0)
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

	dfa := NewDFA("a", "c", states, accept, NewDRuleBook(rules))

	for _, char := range "10101010101101010" {
		dfa.ReadCharacter(string(char))
	}

	if dfa.Rejecting() {
		fmt.Println("This can't be seen !!!")
	}

	fmt.Println(dfa.GetCurrentState()) // Prints the charater 'b'

	// dfa.ReadCharacter("2")
	// if dfa.Rejecting() {
	// 	fmt.Println("This is going to be shown")
	// }
	return dfa
}

func ExampleNFA() *NFA {
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

	if nfa.Accepting() {
		fmt.Println(nfa.GetCurrentState())
	}

	return nfa
}

func ExampleOperations() {
	dfa := x()
	nfa := y()

	x := Union(nfa, dfa)
	y := Concatenation(nfa, dfa)
	z1 := PositiveClosure(dfa)
	z2 := PositiveClosure(nfa)

	deterministic := Determinise(nfa)
	minimal := Minimise(deterministic)

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z1)
	fmt.Println(z2)
	fmt.Println(deterministic)
	fmt.Println(minimal)
}

func ExampleDPDA() {
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

	states := SetFromSlice([]string{"a", "b"})
	accept := SetFromSlice([]string{"b"})

	stack := NewStack()
	stack.Push("#")

	dpda := NewDPDA("a", "b", "a", states, accept, stack, book)

	for _, char := range "000101010101010" {
		fmt.Println(dpda.stack)
		dpda.ReadCharacter(string(char))
	}
	fmt.Println(dpda.stack)
}

func ExampleTuring(t *testing.T) {
	rule1 := NewTRule("alpha", "0", "alpha", "1", "RIGHT")
	rule2 := NewTRule("alpha", "1", "alpha", "0", "RIGHT")
	rule3 := NewTRule("alpha", "#", "sigma", " ", "NOP")

	book := NewEmptyTuringRoolBook()
	book.AddRule(rule1)
	book.AddRule(rule2)
	book.AddRule(rule3)

	accept := NewSet()
	accept.Add("sigma")

	states := NewSet()
	states.Add("alpha")
	states.Add("sigma")
	states.Add("teta")

	tape := NewNonEmptyTape("", "0", "01010001001101110#")

	turing := NewDTuringMachine("alpha", "teta", states, accept, tape, book)

	turing.Run()
	if turing.Rejecting() {
		t.Error("Should not be rejecting!")
	}
	if !turing.Accepting() {
		t.Error("Should be accepting!")
	}
}
