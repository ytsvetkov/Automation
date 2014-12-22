package Automation

import "errors"

type Rule struct {
	from string
	with string
	to   string
}

func NewRule(from, with, to string) (*Rule, error) {
	if len(with) <= 1 {
		return &Rule{from: from, with: with, to: to}, nil
	}
	return nil, errors.New("Too long transition!")
}

func NewRules(from, with, to string) []*Rule {
	rules := make([]*Rule, 0, len(with))
	for _, symbol := range with {
		rule, _ := NewRule(from, string(symbol), to)
		rules = append(rules, rule)
	}
	return rules
}

func (r *Rule) SetStartingState(label string) {
	r.from = label
}

func (r *Rule) SetTransitionRule(symbol string) bool {
	if len(symbol) >= 2 {
		return false
	}
	r.with = symbol
	return true
}

func (r *Rule) SetEndingState(label string) {
	r.to = label
}

func (r *Rule) GetStartingState() string {
	return r.from
}

func (r *Rule) GetTransitionRule() string {
	return r.with
}

func (r *Rule) GetEndingState() string {
	return r.to
}

func (r *Rule) String() string {
	return "(" + r.from + ")-" + r.with + "-(" + r.to + ")"
}

func (r *Rule) New() *Rule {
	return &Rule{from: r.from, with: r.with, to: r.to}
}
