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

func NewRules(from, with, to string) ([]*Rule, error) {
	rules := make([]*Rule, 0, len(with))
	for _, symbol := range with {
		rule, _ := NewRule(from, string(symbol), to)
		rules = append(rules, rule)
	}
	return rules, nil
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
