package Automation

//					     from       with       pop          to  push,direction
type TuringRoolBook map[string]map[string]map[string]map[string][2]string

func NewEmptyTuringRoolBook() TuringRoolBook {
	return make(TuringRoolBook)
}

func (t TuringRoolBook) AddRule(rule TRule) {
}

func (t TuringRoolBook) AddRules(rules []TRule) {
}

func NewTuringRoolBook(rules []TRule) TuringRoolBook {
	book := NewEmptyTuringRoolBook()
	book.AddRules(rules)
	return book
}

func (t TuringRoolBook) String() string {
	return ""
}

// func (t TuringRoolBook) GetFromState(from string) [][5]string {
// }

// func (t TuringRoolBook) GetRuleEnd(from, with, pop string) (result [][3]string) {
// }

// func (t TuringRoolBook) GetFromTransition(from string) (set Set) {
// }

// func (t TuringRoolBook) GetAllRules() (rule []TRule) {
// }
