package Automation

import "errors"

// Imlements the 'RuleBook' interface
// Used for stroring the transition rules of a Deterministic
// Finite Automata. The 'book' has the following structure:
// fromState-withReadCharacter-toState
type DRuleBook map[string]map[string]string

// Empty roolbook
func NewEmptyDRuleBook() DRuleBook {
	return make(DRuleBook)
}

// Adds a single rule iff it does not introduce
// non-deterministic behaviour.
func (d DRuleBook) AddRule(rule *Rule) error {
	if transitons, ok := d[rule.GetFrom()]; ok == false {
		d[rule.GetFrom()] = map[string]string{rule.GetWith(): rule.GetTo()}
	} else if _, okk := transitons[rule.GetWith()]; okk == false {
		d[rule.GetFrom()][rule.GetWith()] = rule.GetTo()
	} else if transitons[rule.GetWith()] != rule.GetTo() {
		return errors.New("This introduces non-deterministic behaviour: " + rule.String() + " !!!")
	}
	return nil
}

// Adds multiple rules and ignores the ones which
// introduce non-deterministic behaviour. As such,
// the order of the rules in the slice is important.
func (d DRuleBook) AddRules(rules []*Rule) error {
	errMsg := "The following rules were not added becaouse of introduction of non-deterministic behaviour: \n"
	var err error
	var flag bool

	for _, rule := range rules {
		err = d.AddRule(rule)
		if err != nil {
			flag = true
			errMsg += rule.String() + "\n"
		}
	}

	if flag {
		return errors.New(errMsg)
	}
	return nil
}

// Returns initialised roolebook. The order of the rules
// in the slice is important. Check documentation for 'AddRules'
func NewDRuleBook(rules []*Rule) (DRuleBook, error) {
	book := NewEmptyDRuleBook()
	err := book.AddRules(rules)
	return book, err
}

func (d DRuleBook) String() string {
	str := "[\n"
	for from, transitons := range d {
		for with, to := range transitons {
			str += "\t" + NewRule(from, with, to).String() + "\n"
		}
	}
	return str + "]"
}

// Returns a slice of the posible transitions from the given
// state. Because this is a deterministic machine, there is
// going to be only one element in it, or none. Always !
func (d DRuleBook) GetFromState(from string) [][2]string {
	tran := make([][2]string, 0, 16)
	if transitons, ok := d[from]; ok != false {
		for with, to := range transitons {
			tran = append(tran, [2]string{with, to})
		}
	}
	return tran
}

// Returns the set of posible states, which are reachable
// with the given transition state. Because this is a
// deterministic machine, there is going to be only one
// element in it, or none. Always !
func (d DRuleBook) GetFromTransition(from, with string) Set {
	if tran, ok := d[from]; ok != false {
		if end, okk := tran[with]; okk != false {
			set := NewSet()
			set.Add(end)
			return set
		}
	}
	return NewSet()
}

// Returns the set of posible states, which are reachable
// with the given transition state. Because this is a
// deterministic machine, there is going to be only one
// element in it, or none. Always !
func (d DRuleBook) GetRuleEnd(from, with string) Set {
	set := NewSet()
	if tran, ok := d[from]; ok != false {
		if end, okk := tran[with]; okk != false {
			set.Add(end)
		}
	}
	return set
}

// Return a slice with all the rules in the curent roolbook
func (d DRuleBook) GetAllRules() (rules []*Rule) {
	for from, tran := range d {
		for with, to := range tran {
			rules = append(rules, NewRule(from, with, to))
		}
	}
	return
}
