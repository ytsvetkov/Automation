package Automation

//					     from       with  to-writer-dir
type TuringRoolBook map[string]map[string][3]string

func NewEmptyTuringRoolBook() TuringRoolBook {
	return make(TuringRoolBook)
}

func (t TuringRoolBook) AddRule(rule TRule) {
	if _, o := t[rule.GetFrom()]; o == false {
		map4 := map[string][3]string{rule.GetWith(): [3]string{rule.GetTo(), rule.GetWriter(), rule.GetDirection()}}
		t[rule.GetFrom()] = map4
	} else if _, okk := t[rule.GetFrom()][rule.GetWith()]; okk == false {
		t[rule.GetFrom()][rule.GetWith()] = [3]string{rule.GetTo(), rule.GetWriter(), rule.GetDirection()}
	}
}

func (t TuringRoolBook) AddRules(rules []TRule) {
	for _, rule := range rules {
		t.AddRule(rule)
	}
}

func NewTuringRoolBook(rules []TRule) TuringRoolBook {
	book := NewEmptyTuringRoolBook()
	book.AddRules(rules)
	return book
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

func (t TuringRoolBook) GetFromState(from string) [][4]string {
	tran := make([][4]string, 0, 16)
	if transitons, ok := t[from]; ok != false {
		for with, triplet := range transitons {
			tran = append(tran, [4]string{with, triplet[0], triplet[1], triplet[2]})
		}
	}
	return tran
}

func (t TuringRoolBook) GetRuleEnd(from, with string) (result [][3]string) {
	if transitons, ok := t[from]; ok != false {
		if triplet, ok := transitons[with]; ok != false {
			result = append(result, [3]string{triplet[0], triplet[1], triplet[2]})
		}
	}
	return
}

func (t TuringRoolBook) GetFromTransition(from string) (set Set) {
	if transitons, ok := t[from]; ok != false {
		for _, triplet := range transitons {
			set.Add(triplet[0])
		}
	}
	return
}

func (t TuringRoolBook) GetAllRules() (rule []TRule) {
	for from, j := range t {
		for with, triplet := range j {
			rule = append(rule, NewTRule(from, with, triplet[0], triplet[1], triplet[2]))
		}
	}
	return
}
