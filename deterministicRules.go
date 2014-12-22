package Automation

import "errors"

type DeterministicRuleBook map[string]map[string]string

func NewDeterministicRuleBook(rules []*Rule) (DeterministicRuleBook, error) {
	book := make(DeterministicRuleBook)

	for _, rule := range rules {
		mapTranEnd, ok := book[rule.GetStartingState()]
		if ok == false {
			book[rule.GetStartingState()] = make(map[string]string)
			book[rule.GetStartingState()][rule.GetTransitionRule()] = rule.GetEndingState()
		} else if _, okk := mapTranEnd[rule.GetTransitionRule()]; okk == false {
			mapTranEnd[rule.GetTransitionRule()] = rule.GetEndingState()
		} else if end, okkk := mapTranEnd[rule.GetTransitionRule()]; okkk == true && end == rule.GetEndingState() {
		} else {
			return nil, errors.New("This set of rules induses non-deterministic behaviour!")
		}
	}

	return book, nil
}

func (r DeterministicRuleBook) String() string {
	str := "[\n"
	for from, rule := range r {
		for i, j := range rule {
			str += "\t(" + from + ")-" + i + "-(" + j + ")\n"
		}
	}
	return str + "]"
}

func (r DeterministicRuleBook) AddRule(rule *Rule) error {
	mapTranEnd, ok := r[rule.GetStartingState()]
	if ok == false {
		r[rule.GetStartingState()] = make(map[string]string)
		r[rule.GetStartingState()][rule.GetTransitionRule()] = rule.GetEndingState()
	} else if _, okk := mapTranEnd[rule.GetTransitionRule()]; okk == false {
		mapTranEnd[rule.GetTransitionRule()] = rule.GetEndingState()
	} else if end, okkk := mapTranEnd[rule.GetTransitionRule()]; okkk == true && end == rule.GetEndingState() {
	} else {
		return errors.New("This rule adds non-deterministic behaviour!")
	}
	return nil
}

func (r DeterministicRuleBook) GetRule(from, with string) *Rule {
	if mapTranEnd, ok := r[from]; ok == false {
		return nil
	} else if end, okk := mapTranEnd[with]; okk == false {
		return nil
	} else {
		rule, _ := NewRule(from, with, end)
		return rule
	}
}
