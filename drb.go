package Automation

type DRuleBook map[string]map[string]string

func NewEmptyDRuleBook() DRuleBook {
	return make(DRuleBook)
}

func (d DRuleBook) AddRule(rule Rule) {
	if transitons, ok := d[rule.GetFrom()]; ok == false {
		d[rule.GetFrom()] = map[string]string{rule.GetWith(): rule.GetTo()}
	} else if _, okk := transitons[rule.GetWith()]; okk == false {
		d[rule.GetFrom()][rule.GetWith()] = rule.GetTo()
	}
}

func (d DRuleBook) AddRules(rules []Rule) {
	for _, rule := range rules {
		d.AddRule(rule)
	}
}

func NewDRuleBook(rules []Rule) DRuleBook {
	book := NewEmptyDRuleBook()
	book.AddRules(rules)
	return book
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

func (d DRuleBook) GetFromState(from string) [][2]string {
	tran := make([][2]string, 0, 16)
	if transitons, ok := d[from]; ok != false {
		for with, to := range transitons {
			tran = append(tran, [2]string{with, to})
		}
	}
	return tran
}

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

func (d DRuleBook) GetRuleEnd(from, with string) Set {
	set := NewSet()
	if tran, ok := d[from]; ok != false {
		if end, okk := tran[with]; okk != false {
			set.Add(end)
		}
	}
	return set
}

func (d DRuleBook) GetAllRules() (rules []Rule) {
	for from, tran := range d {
		for with, to := range tran {
			rules = append(rules, NewRule(from, with, to))
		}
	}
	return
}
