package Automation

// Imlements the 'RuleBook' interface
// Used for stroring the transition rules of a NonDeterministic
// Finite Automata. The 'book' has the following structure:
// fromState-withReadCharacter-toState
type NRuleBook map[string]map[string]Set

func NewEmptyNRuleBook() NRuleBook {
	return make(NRuleBook)
}

// Adds a single rule and allways succeeds.
func (n NRuleBook) AddRule(rule *Rule) error {
	if tran, ok := n[rule.GetFrom()]; ok == false {
		set := NewSet()
		set.Add(rule.GetTo())
		n[rule.GetFrom()] = map[string]Set{rule.GetWith(): set}
	} else if _, okk := tran[rule.GetWith()]; okk == false {
		set := NewSet()
		set.Add(rule.GetTo())
		n[rule.GetFrom()][rule.GetWith()] = set
	} else {
		n[rule.GetFrom()][rule.GetWith()].Add(rule.GetTo())
	}
	return nil
}

// Adds multiple rules and allways succeeds. As such,
// the order of the rules in the slice is not important.
func (n NRuleBook) AddRules(rules []*Rule) error {
	for _, rule := range rules {
		n.AddRule(rule)
	}
	return nil
}

// Returns initialised roolebook.
func NewNRuleBook(rules []*Rule) NRuleBook {
	book := NewEmptyNRuleBook()
	book.AddRules(rules)
	return book
}

func (n NRuleBook) String() string {
	str := "[\n"
	for from, transitons := range n {
		for with, set := range transitons {
			for to, _ := range set {
				str += "\t" + NewRule(from, with, to).String() + "\n"
			}
		}
	}
	return str + "]"
}

// Returns a slice of the posible transitions from the given state.
func (n NRuleBook) GetFromState(from string) [][2]string {
	tran := make([][2]string, 0, 16)
	if transitons, ok := n[from]; ok != false {
		for to, ends := range transitons {
			for _, end := range ends.Values() {
				tran = append(tran, [2]string{to, end})
			}
		}
	}
	return tran
}

// Returns the set of posible states, which are reachable
// with the given transition state.
func (n NRuleBook) GetFromTransition(from, with string) Set {
	if tran, ok := n[from]; ok != false {
		if end, okk := tran[with]; okk != false {
			return end
		}
	}
	return NewSet()
}

// Returns the set of posible states, which are reachable
// from the given transition state.
func (n NRuleBook) GetRuleEnd(from, with string) (set Set) {
	if tran, ok := n[from]; ok != false {
		if end, okk := tran[with]; okk != false {
			return end
		}
	}
	return set
}

// Return a slice with all the rules in the curent roolbook
func (n NRuleBook) GetAllRules() (rules []*Rule) {
	for from, tran := range n {
		for with, tos := range tran {
			for _, to := range tos.Values() {
				rules = append(rules, NewRule(from, with, to))
			}
		}
	}
	return
}
