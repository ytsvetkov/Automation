package Automation

type TRule struct {
	with      string
	writer    string
	push      string
	pop       string
	from      string
	to        string
	direction string
}

func NewTRule(from, with, pop, to, push, direction string) TRule {
	return TRule{from: from, with: with, to: to, pop: pop, push: push, direction: direction}
}

func NewTRules(from, with, pop, to, push, direction string) []TRule {
	rules := make([]TRule, 0, len(with))
	for _, char := range with {
		rules = append(rules, NewTRule(from, string(char), pop, to, push, direction))
	}
	return rules
}

func (t TRule) String() string {
	return "(" + t.from + ")-" + t.with + "-|*" + t.pop + "*|" + "(" + t.to + ")<" + t.push + ">" + "[" + t.direction + "]"
}
