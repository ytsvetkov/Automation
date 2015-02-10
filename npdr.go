package Automation

//					 from		with        pop       push  to
type NPRuleBook map[string]map[string]map[string]map[string]Set

func NewEmptyNPRuleBook() NPRuleBook {
	return make(NPRuleBook)
}

func (d NPRuleBook) AddRule(rule *PRule) {
	// ToDo
	// ToDo
	// ToDo ToDo ToDo ToDooo
	// ToDoDoDo
}

func (d NPRuleBook) AddRules(rules []*PRule) {
	for _, rule := range rules {
		d.AddRule(rule)
	}
}

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

func (d NPRuleBook) GetRuleEnd(from, with, pop string) (result [][2]string) {
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
	return
}

func (d NPRuleBook) GetFromTransition(from string) (set Set) {
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
	return
}

func (d NPRuleBook) GetAllRules() (rule []*PRule) {
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
	return
}
