package Automation

type NonDeterministicRuleBook map[string]map[string]*Set

func NewNonDeterministicRuleBook(rules []*Rule) NonDeterministicRuleBook {
	book := make(NonDeterministicRuleBook)

	for _, rule := range rules {
		mapTranEnd, ok := book[rule.GetStartingState()]
		if ok == false {
			set := NewSet()
			set.Add(rule.GetEndingState())
			book[rule.GetStartingState()] = map[string]*Set{rule.GetTransitionRule(): set}
		} else if end, okk := mapTranEnd[rule.GetTransitionRule()]; okk == false {
			set := NewSet()
			set.Add(rule.GetEndingState())
			mapTranEnd[rule.GetTransitionRule()] = set
		} else {
			end.Add(rule.GetEndingState())
		}

	}
	return book
}

func (r NonDeterministicRuleBook) String() string {
	str := "[\n"
	for from, mapTranEnd := range r {
		for tran, set := range mapTranEnd {
			for _, member := range set.Values() {
				str += "\t(" + from + ")-" + tran + "-(" + member + ")\n"
			}
		}
	}
	return str + "]"
}

func (r NonDeterministicRuleBook) AddRule(rule *Rule) {
	mapTranEnd, ok := r[rule.GetStartingState()]
	if ok == false {
		set := NewSet()
		set.Add(rule.GetEndingState())
		r[rule.GetStartingState()] = map[string]*Set{rule.GetTransitionRule(): set}
	} else if end, okk := mapTranEnd[rule.GetTransitionRule()]; okk == false {
		set := NewSet()
		set.Add(rule.GetEndingState())
		mapTranEnd[rule.GetTransitionRule()] = set
	} else {
		end.Add(rule.GetEndingState())
	}
}

func (r NonDeterministicRuleBook) GetRulesEnd(from, with string) Set {
	if mapTranEnd, ok := r[from]; ok == false {
		return nil
	} else if end, okk := mapTranEnd[with]; okk == false {
		return nil
	} else {
		return *end
	}
}
