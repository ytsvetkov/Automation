package Automation

type Rule struct {
	from string
	with string
	to   string
}

func NewRule(from, with, to string) Rule {
	return Rule{from: from, with: with, to: to}
}

func NewRules(from, with, to string) []Rule {
	rules := make([]Rule, 0, len(with))
	for _, char := range with {
		rules = append(rules, NewRule(from, string(char), to))
	}
	return rules
}
func (r Rule) String() string {
	return "(" + r.from + ")-" + r.with + "-(" + r.to + ")"
}

func (r Rule) GetFrom() string {
	return r.from
}

func (r Rule) GetWith() string {
	return r.with
}

func (r Rule) GetTo() string {
	return r.to
}

func (r Rule) SetFrom(from string) {
	r.from = from
}

func (r Rule) SetWith(with string) {
	if len(with) <= 1 {
		r.with = with
	}
}

func (r Rule) SetTo(to string) {
	r.to = to
}
