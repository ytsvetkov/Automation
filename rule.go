package Automation

import "errors"

// Rules( or transition rules) for the (D|N)FA.
type Rule struct {
	from string
	with string
	to   string
}

func NewRule(from, with, to string) *Rule {
	if len(with) == 1 {
		return &Rule{from: from, with: with, to: to}
	}
	return nil
}

func NewRules(from, with, to string) []*Rule {
	rules := make([]*Rule, 0, len(with))
	for _, char := range with {
		rules = append(rules, NewRule(from, string(char), to))
	}
	return rules
}
func (r *Rule) String() string {
	return "(" + r.from + ")-" + r.with + "-(" + r.to + ")"
}

func (r *Rule) GetFrom() string {
	return r.from
}

func (r *Rule) GetWith() string {
	return r.with
}

func (r *Rule) GetTo() string {
	return r.to
}

func (r *Rule) SetFrom(from string) {
	r.from = from
}

func (r *Rule) SetWith(with string) error {
	if len(with) == 1 {
		r.with = with
		return nil
	}
	return errors.New("Must be a single letter!")
}

func (r *Rule) SetTo(to string) {
	r.to = to
}
