package Automation

func Union(auto1, auto2 RegularAutomata) NFA {
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

func Concatenation(auto1, auto2 RegularAutomata) NFA {
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

func PositiveClosure(auto RegularAutomata) NFA {
	rules := []Rule{}
	for _, acc := range auto.GetAcceptStates().Values() {
		for _, start := range auto.GetStartStates().Values() {
			for _, tran := range auto.GetFromState(start) {
				rules = append(rules, NewRule(acc, tran[0], tran[1]))
			}
		}
	}
	return NewNFA(auto.GetStartStates(), auto.GetAllStates(), auto.GetAcceptStates(), auto.GetReject(), NewNRuleBook(rules))
}

// func Determinise(nfa NFA) DFA {

// }

// func Minimise(dfa DFA) DFA {

// }
