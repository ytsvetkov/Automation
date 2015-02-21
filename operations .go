package Automation

// Union of two finite automata. This generates new automata,
// whos recognised language is the union of the languages of
// the two given automata.
func Union(auto1, auto2 RegularAutomata) *NFA {
	start := auto1.GetStartStates()
	start.AddSet(auto2.GetStartStates())

	states := NewSet()
	states.AddSet(auto1.GetAllStates())
	states.AddSet(auto2.GetAllStates())
	states.AddSet(start)

	accept := NewSet()
	if auto1.GetStartStates().Intersection(auto1.GetAcceptStates()).Cardinality() != 0 || auto2.GetStartStates().Intersection(auto2.GetAcceptStates()).Cardinality() != 0 {
		accept.AddSet(start)
	}
	accept.AddSet(auto1.GetAcceptStates())
	accept.AddSet(auto2.GetAcceptStates())

	reject := auto1.GetReject() + "+" + auto2.GetReject()

	rules := auto1.GetAllRules()
	rules = append(rules, auto2.GetAllRules()...)
	for _, state := range auto1.GetAllStates().Values() {
		for _, tran := range auto1.GetFromState(state) {
			for _, starting := range start.Values() {
				rules = append(rules, NewRule(starting, tran[0], tran[1]))
			}
		}
	}
	for _, state := range auto2.GetAllStates().Values() {
		for _, tran := range auto2.GetFromState(state) {
			for _, starting := range start.Values() {
				rules = append(rules, NewRule(starting, tran[0], tran[1]))
			}
		}
	}
	return NewNFA(start, states, accept, reject, NewNRuleBook(rules))
}

// Concatenation of two finite automata. This generates new automata,
// whos recognised language is the concatenation of the languages of
// the two given automata.
func Concatenation(auto1, auto2 RegularAutomata) *NFA {
	start := auto1.GetStartStates()

	states := NewSet()
	states.AddSet(auto1.GetAllStates())
	states.AddSet(auto2.GetAllStates())

	accept := NewSet()
	if auto2.GetStartStates().Intersection(auto2.GetAcceptStates()).Cardinality() != 0 {
		accept.AddSet(auto1.GetStartStates())
	}
	accept.AddSet(auto2.GetAcceptStates())

	reject := auto1.GetReject() + "+" + auto2.GetReject()

	rules := auto2.GetAllRules()
	rules = append(rules, auto1.GetAllRules()...)
	for _, astate := range auto1.GetAcceptStates().Values() {
		for _, sstate := range auto2.GetStartStates().Values() {
			for _, tran := range auto2.GetFromState(sstate) {
				rules = append(rules, NewRule(astate, tran[0], tran[1]))
			}
		}
	}
	return NewNFA(start, states, accept, reject, NewNRuleBook(rules))
}

// Posotive closure of finite automata. If 'auto' can recognise
// the language 'L', PositiveClosure(auto) recognises L*
func PositiveClosure(auto RegularAutomata) *NFA {
	rules := []*Rule{}
	for _, acc := range auto.GetAcceptStates().Values() {
		for _, start := range auto.GetStartStates().Values() {
			for _, tran := range auto.GetFromState(start) {
				rules = append(rules, NewRule(acc, tran[0], tran[1]))
			}
		}
	}
	return NewNFA(auto.GetStartStates(), auto.GetAllStates(), auto.GetAcceptStates(), auto.GetReject(), NewNRuleBook(rules))
}

// Powerset construction to get determinised automata,
// recognising the same language as the given NFA
func Determinise(nfa *NFA) *DFA {
	startStates := nfa.GetStartStates()
	start := startStates.String()

	accept := NewSet()
	states := NewSet()

	alphabet := nfa.GetAlphabet().Values()
	setStates := []Set{startStates}
	finalised := []Set{}
	table := make(map[string]map[string]Set)

	for len(setStates) > 0 {
		newSetState := setStates[0]

		if len(setStates) > 1 {
			setStates = setStates[1 : len(setStates)-1]
		} else {
			setStates = []Set{}
		}

		finalised = append(finalised, newSetState)

		tran := make(map[string]Set)
		for _, letter := range alphabet {
			tran[letter] = NewSet()
		}

		for _, letter := range alphabet {
			for _, state := range newSetState.Values() {
				x := nfa.rules.GetRuleEnd(state, letter)
				tran[letter].AddSet(x)

				if nfa.GetAcceptStates().Intersection(x).Cardinality() > 0 {
					accept.AddSet(x)
				}
			}
		}

		table[newSetState.String()] = tran
		states.Add(newSetState.String())

		for _, set := range tran {
			if set.Cardinality() == 0 {
				goto skip
			}

			for _, set2 := range finalised {
				if set.Eq(set2) {
					goto skip
				}
			}

			for _, set2 := range setStates {
				if set.Eq(set2) {
					goto skip
				}
			}

			setStates = append(setStates, set)
		skip:
		}
	}

	book := NewEmptyDRuleBook()
	for from, trans := range table {
		for with, to := range trans {
			if to.Cardinality() == 0 {
				continue
			}
			book.AddRule(NewRule(from, with, to.String()))
		}
	}

	return NewDFA(start, nfa.GetReject(), states, accept, book)

}

// Function for minimisation of deterministic
// finite automata. It uses the Brzozowski's algorithm
// which is suitable for some "small" automatas and
// for the majority of the "larger" ones.
func Minimise(dfa *DFA) *DFA {
	nfa1 := ReverseAutomata(dfa)
	dfa2 := Determinise(nfa1)
	nfa3 := ReverseAutomata(dfa2)
	return Determinise(nfa3)
}

// Reversing all the edges, making the initial state an accept state,
// and the accept states initial, to get an NFA for the reverse language.
func ReverseAutomata(dfa *DFA) *NFA {
	accept := dfa.GetStartStates()
	start := dfa.GetAcceptStates()
	reject := dfa.reject

	rules := dfa.GetAllRules()
	book := NewEmptyNRuleBook()
	for _, rule := range rules {
		book.AddRule(NewRule(rule.GetTo(), rule.GetWith(), rule.GetFrom()))
	}

	return NewNFA(start, dfa.GetAllStates(), accept, reject, book)
}
