package Automation

import "errors"

// Because of this: https://groups.google.com/forum/#!topic/golang-nuts/VUtUmxm2ubU
//					from		with        pop        to    push
type DPRuleBook map[string]map[string]map[string]map[string]string

func NewEmptyDPRuleBook() DPRuleBook {
	return make(DPRuleBook)
}

//Adds a single rule iff it does not introduce
//non-deterministic behaviour.
func (d DPRuleBook) AddRule(rule *PRule) error {
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
	} else if _, ko := d[rule.GetFrom()][rule.GetWith()][rule.GetPop()]; ko == true {
		return errors.New("This introduces non-deterministic behaviour: " + rule.String() + " !!!")
	} else if _, okkk := d[rule.GetFrom()][rule.GetWith()][rule.GetPop()][rule.GetTo()]; okkk == false {
		d[rule.GetFrom()][rule.GetWith()][rule.GetPop()][rule.GetTo()] = rule.GetPush()
	} else if d[rule.GetFrom()][rule.GetWith()][rule.GetPop()][rule.GetTo()] != rule.GetPush() {
		return errors.New("This introduces non-deterministic behaviour: " + rule.String() + " !!!")
	}
	return nil
}

//Adds multiple rules and ignores the ones which
//introduce non-deterministic behaviour. As such,
//the order of the rules in the slice is important.
func (d DPRuleBook) AddRules(rules []*PRule) error {
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

func NewDPRuleBook(rules []*PRule) (DPRuleBook, error) {
	book := NewEmptyDPRuleBook()
	err := book.AddRules(rules)
	return book, err
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

//Returns the posible transitions from the given state.
//Each element in the slice is of the form:
//with - pop - to - push
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

//Returns slice of posible state-push tuples, which are
//reachable with the given transition state. Because
//this is a deterministic machine, there is going to
//be only one element in it, or none. Always !
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

//Returns the set of all reachable states from the given one.
func (d DPRuleBook) GetFromTransition(from string) Set {
	set := NewSet()
	if transitons, ok := d[from]; ok != false {
		for _, rest := range transitons {
			for _, rest2 := range rest {
				for to, _ := range rest2 {
					set.Add(to)
				}
			}
		}
	}
	return set
}

//Return a slice with all the rules in the curent roolbook
func (d DPRuleBook) GetAllRules() (rule []*PRule) {
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
