package Automation

import "errors"

//					     from       with  to-writer-dir
type TuringRoolBook map[string]map[string][3]string

func NewEmptyTuringRoolBook() TuringRoolBook {
	return make(TuringRoolBook)
}

//Adds a single rule iff it does not introduce
//non-deterministic behaviour.
func (t TuringRoolBook) AddRule(rule *TRule) error {
	if _, o := t[rule.GetFrom()]; o == false {
		map4 := map[string][3]string{rule.GetWith(): [3]string{rule.GetTo(), rule.GetWriter(), rule.GetDirection()}}
		t[rule.GetFrom()] = map4
	} else if _, okk := t[rule.GetFrom()][rule.GetWith()]; okk == false {
		t[rule.GetFrom()][rule.GetWith()] = [3]string{rule.GetTo(), rule.GetWriter(), rule.GetDirection()}
	} else if t[rule.GetFrom()][rule.GetWith()] == [3]string{rule.GetTo(), rule.GetWriter(), rule.GetDirection()} {
		return errors.New("This introduces non-deterministic behaviour: " + rule.String() + " !!!")
	}
	return nil
}

//Adds multiple rules and ignores the ones which
//introduce non-deterministic behaviour. As such,
//the order of the rules in the slice is important.
func (t TuringRoolBook) AddRules(rules []*TRule) error {
	errMsg := "The following rules were not added becaouse of introduction of non-deterministic behaviour: \n"
	var err error
	var flag bool

	for _, rule := range rules {
		t.AddRule(rule)
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

func NewTuringRoolBook(rules []*TRule) (TuringRoolBook, error) {
	book := NewEmptyTuringRoolBook()
	err := book.AddRules(rules)
	return book, err
}

func (t TuringRoolBook) String() string {
	str := "[\n"
	for from, j := range t {
		for with, triplet := range j {
			str += "\t" + NewTRule(from, with, triplet[0], triplet[1], triplet[2]).String() + "\n"
		}
	}
	return str
}

//Returns the posible transitions from the given state.
//Each element in the slice is of the form:
//with - to - write - dir
func (t TuringRoolBook) GetFromState(from string) [][4]string {
	tran := make([][4]string, 0, 16)
	if transitons, ok := t[from]; ok != false {
		for with, triplet := range transitons {
			tran = append(tran, [4]string{with, triplet[0], triplet[1], triplet[2]})
		}
	}
	return tran
}

//Returns slice of posible state-push tuples, which are
//reachable with the given transition state. Because
//this is a deterministic machine, there is going to
//be only one element in it, or none. Always !
func (t TuringRoolBook) GetRuleEnd(from, with string) [][3]string {
	result := [][3]string{}
	if transitons, ok := t[from]; ok != false {
		if triplet, ok := transitons[with]; ok != false {
			result = append(result, [3]string{triplet[0], triplet[1], triplet[2]})
		}
	}
	return result
}

//Returns the set of posible states, which are
//reachable with the given transition state.
func (t TuringRoolBook) GetFromTransition(from string) Set {
	set := NewSet()
	if transitons, ok := t[from]; ok != false {
		for _, triplet := range transitons {
			set.Add(triplet[0])
		}
	}
	return set
}

//Return a slice with all the rules in the curent roolbook.
func (t TuringRoolBook) GetAllRules() []*TRule {
	rule := []*TRule{}
	for from, j := range t {
		for with, triplet := range j {
			rule = append(rule, NewTRule(from, with, triplet[0], triplet[1], triplet[2]))
		}
	}
	return rule
}
