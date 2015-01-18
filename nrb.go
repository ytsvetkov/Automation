package Automation

type NRuleBook map[string]map[string]Set

func NewEmptyNRuleBook() NRuleBook {
	return make(NRuleBook)
}

func (n NRuleBook) AddRule(rule Rule) {
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
}

func (n NRuleBook) AddRules(rules []Rule) {
	for _, rule := range rules {
		n.AddRule(rule)
	}
}

func NewNRuleBook(rules []Rule) NRuleBook {
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

func (n NRuleBook) GetFromTransition(from, with string) Set {
	if tran, ok := n[from]; ok != false {
		if end, okk := tran[with]; okk != false {
			return end
		}
	}
	return NewSet()
}

func (n NRuleBook) GetRuleEnd(from, with string) (set Set) {
	if tran, ok := n[from]; ok != false {
		if end, okk := tran[with]; okk != false {
			return end
		}
	}
	return set
}

func (n NRuleBook) GetAllRules() (rules []Rule) {
	for from, tran := range n {
		for with, tos := range tran {
			for _, to := range tos.Values() {
				rules = append(rules, NewRule(from, with, to))
			}
		}
	}
	return
}
