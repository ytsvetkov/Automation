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

type PRule struct {
	from string
	with string
	to   string
	pop  string
	push string
}

func NewPRule(from, with, pop, to, push string) PRule {
	return PRule{from: from, with: with, to: to, pop: pop, push: push}
}

func NewPRules(from, with, to, pop, push string) []PRule {
	rules := make([]PRule, 0, len(with))
	for _, char := range with {
		rules = append(rules, NewPRule(from, string(char), to, push, pop))
	}
	return rules
}

func (p PRule) String() string {
	return "(" + p.from + ")-" + p.with + "-|*" + p.pop + "*|" + "(" + p.to + ")<" + p.push + ">"
}

func (p PRule) GetFrom() string {
	return p.from
}

func (p PRule) GetWith() string {
	return p.with
}

func (p PRule) GetTo() string {
	return p.to
}

func (p PRule) GetPush() string {
	return p.push
}

func (p PRule) GetPop() string {
	return p.pop
}

func (p PRule) SetFrom(from string) {
	p.from = from
}

func (p PRule) SetWith(with string) {
	if len(with) <= 1 {
		p.with = with
	}
}

func (p PRule) SetTo(to string) {
	p.to = to
}

func (p PRule) SetPush(push string) {
	p.push = push
}
func (p PRule) SetPop(pop string) {
	p.pop = pop
}
