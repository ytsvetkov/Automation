package Automation

//					from		with        pop        to    push
type DPRuleBook map[string]map[string]map[string]map[string]string

func NewEmptyDPRuleBook() DPRuleBook {
	return make(DPRuleBook)
}

func (d DPRuleBook) AddRule(rule PRule) {
	if _, o := d[rule.GetFrom()]; o == false {
		map4 := map[string]string{rule.GetTo(): rule.GetPush()}
		map3 := map[string]map[string]string{rule.GetPop(): map4}
		map2 := map[string]map[string]map[string]string{rule.GetWith(): map3}
		d[rule.GetFrom()] = map2
	} else if _, ok := d[rule.GetFrom()][rule.GetWith()]; ok == false {
		map3 := map[string]string{rule.GetTo(): rule.GetPush()}
		map2 := map[string]map[string]string{rule.GetPop(): map3}
		d[rule.GetFrom()][rule.GetWith()] = map2
	} else if _, okk := d[rule.GetFrom()][rule.GetWith()][rule.GetPop()]; okk == false {
		map3 := map[string]string{rule.GetTo(): rule.GetPush()}
		d[rule.GetFrom()][rule.GetWith()][rule.GetPop()] = map3
	} else if _, okkk := d[rule.GetFrom()][rule.GetWith()][rule.GetPop()][rule.GetTo()]; okkk == false {
		d[rule.GetFrom()][rule.GetWith()][rule.GetPop()][rule.GetTo()] = rule.GetPush()
	}
}

func (d DPRuleBook) AddRules(rules []PRule) {
	for _, rule := range rules {
		d.AddRule(rule)
	}
}

func NewDPRuleBook(rules []PRule) DPRuleBook {
	book := NewEmptyDPRuleBook()
	book.AddRules(rules)
	return book
}

func (d DPRuleBook) String() string {
	str := "[\n"
	for from, rest := range d {
		for with, rest2 := range rest {
			for pop, rest3 := range rest2 {
				for to, push := range rest3 {
					str += "\t" + NewPRule(from, with, pop, to, push).String() + "\n"
				}
			}
		}
	}
	return str + "]"
}

func (d DPRuleBook) GetFromState(from string) [][4]string {
	tran := make([][4]string, 0, 16)
	if transitons, ok := d[from]; ok != false {
		for with, rest := range transitons {
			for pop, rest2 := range rest {
				for to, push := range rest2 {
					tran = append(tran, [4]string{with, to, pop, push})
				}
			}
		}
	}
	return tran
}

func (d DPRuleBook) GetRuleEnd(from, with, pop string) (result [][2]string) {
	if transitons, ok := d[from]; ok != false {
		if transitons2, okk := transitons[with]; okk != false {
			if transitons3, okkk := transitons2[pop]; okkk != false {
				for to, push := range transitons3 {
					result = append(result, [2]string{to, push})
					return
				}
			}
		}
	}
	return
}

func (d DPRuleBook) GetFromTransition(from string) (set Set) {
	if transitons, ok := d[from]; ok != false {
		for _, rest := range transitons {
			for _, rest2 := range rest {
				for to, _ := range rest2 {
					set.Add(to)
				}
			}
		}
	}
	return
}

func (d DPRuleBook) GetAllRules() (rule []PRule) {
	for from, rest := range d {
		for with, rest2 := range rest {
			for pop, rest3 := range rest2 {
				for to, push := range rest3 {
					rule = append(rule, NewPRule(from, with, pop, to, push))
				}
			}
		}
	}
	return
}
