package Automation

// Imlements the 'RuleBook' interface
// Used for stroring the transition rules of a NonDeterministic
// Pushdown Automata. The 'book' has the following structure:
// fromState-withReadCharacter-popCharacter-pushCharacter-toStates
// Its used this structure because of https://groups.google.com/forum/#!topic/golang-nuts/VUtUmxm2ubU
type NPRuleBook map[string]map[string]map[string]map[string]Set

func NewEmptyNPRuleBook() NPRuleBook {
	return make(NPRuleBook)
}

// Adds a single rule and allways succeeds.
func (d NPRuleBook) AddRule(rule *PRule) error {
	if _, o := d[rule.GetFrom()]; o == false {
		set := NewSet()
		set.Add(rule.GetTo())
		map4 := map[string]Set{rule.GetPush(): set}
		map3 := map[string]map[string]Set{rule.GetPop(): map4}
		map2 := map[string]map[string]map[string]Set{rule.GetWith(): map3}
		d[rule.GetFrom()] = map2
	} else if _, ok := d[rule.GetFrom()][rule.GetWith()]; ok == false {
		set := NewSet()
		set.Add(rule.GetTo())
		map3 := map[string]Set{rule.GetPush(): set}
		map2 := map[string]map[string]Set{rule.GetPop(): map3}
		d[rule.GetFrom()][rule.GetWith()] = map2
	} else if _, okk := d[rule.GetFrom()][rule.GetWith()][rule.GetPop()]; okk == false {
		set := NewSet()
		set.Add(rule.GetTo())
		map3 := map[string]Set{rule.GetPush(): set}
		d[rule.GetFrom()][rule.GetWith()][rule.GetPop()] = map3
	} else if _, okkk := d[rule.GetFrom()][rule.GetWith()][rule.GetPop()][rule.GetPush()]; okkk == false {
		set := NewSet()
		set.Add(rule.GetTo())
		d[rule.GetFrom()][rule.GetWith()][rule.GetPop()][rule.GetPush()] = set
	} else {
		d[rule.GetFrom()][rule.GetWith()][rule.GetPop()][rule.GetPush()].Add(rule.GetPush())
	}
	return nil
}

// Adds multiple rules and allways succeeds. As such,
// the order of the rules in the slice is not important.
func (d NPRuleBook) AddRules(rules []*PRule) error {
	for _, rule := range rules {
		d.AddRule(rule)
	}
	return nil
}

// Returns initialised roolebook.
func NewNPRuleBook(rules []*PRule) NPRuleBook {
	book := NewEmptyNPRuleBook()
	book.AddRules(rules)
	return book
}

func (d NPRuleBook) String() string {
	str := "[\n"
	for from, rest := range d {
		for with, rest2 := range rest {
			for pop, rest3 := range rest2 {
				for push, set := range rest3 {
					for _, to := range set.Values() {
						str += "\t" + NewPRule(from, with, pop, to, push).String() + "\n"
					}
				}
			}
		}
	}
	return str + "]"
}

// Returns a slice of the posible transitions from the given state.
func (d NPRuleBook) GetFromState(from string) [][4]string {
	tran := make([][4]string, 0, 16)
	if transitons, ok := d[from]; ok != false {
		for with, rest := range transitons {
			for pop, rest2 := range rest {
				for push, set := range rest2 {
					for _, to := range set.Values() {
						tran = append(tran, [4]string{with, to, pop, push})
					}
				}
			}
		}
	}
	return tran
}

// Returns a slice of the posible transition-ends from the given state.
func (d NPRuleBook) GetRuleEnd(from, with, pop string) [][2]string {
	result := [][2]string{}
	if transitons, ok := d[from]; ok != false {
		if transitons2, okk := transitons[with]; okk != false {
			if transitons3, okkk := transitons2[pop]; okkk != false {
				for push, set := range transitons3 {
					for _, to := range set.Values() {
						result = append(result, [2]string{to, push})
					}
				}
			}
		}
	}
	return result
}

// Returns the set of posible states, which are reachable
// from the given state.
func (d NPRuleBook) GetFromTransition(from string) Set {
	set := NewSet()
	if transitons, ok := d[from]; ok != false {
		for _, rest := range transitons {
			for _, rest2 := range rest {
				for _, rest3 := range rest2 {
					for _, to := range rest3.Values() {
						set.Add(to)
					}
				}
			}
		}
	}
	return set
}

// Return a slice with all the rules in the curent roolbook
func (d NPRuleBook) GetAllRules() []*PRule {
	rule := []*PRule{}
	for from, rest := range d {
		for with, rest2 := range rest {
			for pop, rest3 := range rest2 {
				for push, set := range rest3 {
					for _, to := range set.Values() {
						rule = append(rule, NewPRule(from, with, pop, to, push))
					}
				}
			}
		}
	}
	return rule
}
