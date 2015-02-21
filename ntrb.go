package Automation

// Imlements the 'RuleBook' interface
// Used for stroring the transition rules of a NonDeterministic
// Turing machine. The 'book' has the following structure:
// fromState-withReadCharacter-toState-writeCharacter-direction
type TuringNRoolBook map[string]map[string][][3]string

func NewEmptyNTuringNRoolBook() TuringNRoolBook {
	return make(TuringNRoolBook)
}

// Adds a single rule.
func (t TuringNRoolBook) AddRule(rule *TRule) error {
	if _, o := t[rule.GetFrom()]; o == false {
		array := [3]string{rule.GetTo(), rule.GetWriter(), rule.GetDirection()}
		slice := [][3]string{array}
		map4 := map[string][][3]string{rule.GetWith(): slice}
		t[rule.GetFrom()] = map4
	} else if _, okk := t[rule.GetFrom()][rule.GetWith()]; okk == false {
		t[rule.GetFrom()][rule.GetWith()] = make([][3]string, 0)
		t[rule.GetFrom()][rule.GetWith()] = append(t[rule.GetFrom()][rule.GetWith()], [3]string{rule.GetTo(), rule.GetWriter(), rule.GetDirection()})
	} else {
		for _, slice := range t[rule.GetFrom()][rule.GetWith()] {
			if slice[0] == rule.GetTo() && slice[1] == rule.GetWriter() && slice[2] == rule.GetDirection() {
				return nil
			}
		}
		t[rule.GetFrom()][rule.GetWith()] = append(t[rule.GetFrom()][rule.GetWith()], [3]string{rule.GetTo(), rule.GetWriter(), rule.GetDirection()})
	}
	return nil
}

// Adds multiple rules . As such, the order
// of the rules in the slice is not important.
func (t TuringNRoolBook) AddRules(rules []*TRule) error {
	for _, rule := range rules {
		t.AddRule(rule)
	}
	return nil
}

func NewTuringNRoolBook(rules []*TRule) (TuringNRoolBook, error) {
	book := NewEmptyNTuringNRoolBook()
	err := book.AddRules(rules)
	return book, err
}

func (t TuringNRoolBook) String() string {
	str := "[\n"
	for from, j := range t {
		for with, slice := range j {
			for _, triplet := range slice {
				str += "\t" + NewTRule(from, with, triplet[0], triplet[1], triplet[2]).String() + "\n"
			}
		}
	}
	return str
}

// Returns the posible transitions from the given state.
// Each element in the slice is of the form:
// with - to - write - dir
func (t TuringNRoolBook) GetFromState(from string) [][4]string {
	tran := make([][4]string, 0, 16)
	if transitons, ok := t[from]; ok != false {
		for with, slice := range transitons {
			for _, triplet := range slice {
				tran = append(tran, [4]string{with, triplet[0], triplet[1], triplet[2]})
			}
		}
	}
	return tran
}

// Returns slice of posible state-push tuples, which are
// reachable with the given transition state.
func (t TuringNRoolBook) GetRuleEnd(from, with string) [][3]string {
	result := make([][3]string, 0)
	if transitons, ok := t[from]; ok != false {
		if slice, ok := transitons[with]; ok != false {
			for _, triplet := range slice {
				result = append(result, [3]string{triplet[0], triplet[1], triplet[2]})
			}
		}
	}
	return result
}

// Returns the set of posible states, which are
// reachable with the given transition state.
func (t TuringNRoolBook) GetFromTransition(from string) Set {
	set := NewSet()
	if transitons, ok := t[from]; ok != false {
		for _, slice := range transitons {
			for _, triplet := range slice {
				set.Add(triplet[0])
			}
		}
	}
	return set
}

// Return a slice with all the rules in the curent roolbook.
func (t TuringNRoolBook) GetAllRules() []*TRule {
	rule := []*TRule{}
	for from, j := range t {
		for with, slice := range j {
			for _, triplet := range slice {
				rule = append(rule, NewTRule(from, with, triplet[0], triplet[1], triplet[2]))
			}
		}
	}
	return rule
}
