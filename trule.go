package Automation

import "errors"

type TRule struct {
	with      string
	writer    string
	from      string
	to        string
	direction string
}

func NewTRule(from, with, to, writer, direction string) *TRule {
	if len(with) != 1 {
		return nil
	} else if (direction != "LEFT") && (direction != "RIGHT") && (direction != "NOP") {
		return nil
	}
	return &TRule{from: from, with: with, to: to, writer: writer, direction: direction}
}

func NewTRules(from, with, to, writer, direction string) []*TRule {
	rules := make([]*TRule, 0, len(with))
	for _, char := range with {
		rules = append(rules, NewTRule(from, string(char), to, writer, direction))
	}
	return rules
}

func (t *TRule) String() string {
	return "(" + t.from + ")-" + t.with + "-(" + t.to + ")<" + t.writer + ">[" + t.direction + "]"
}

func (t *TRule) GetWith() string {
	return t.with
}

func (t *TRule) GetWriter() string {
	return t.writer
}

func (t *TRule) GetTo() string {
	return t.to
}

func (t *TRule) GetFrom() string {
	return t.from
}

func (t *TRule) GetDirection() string {
	return t.direction
}

func (t *TRule) SetFrom(from string) {
	t.from = from
}

func (t *TRule) SetWith(with string) error {
	if len(with) == 1 {
		t.with = with
		return nil
	}
	return errors.New("Must be a single letter!")
}

func (t *TRule) SetTo(to string) {
	t.to = to
}

func (t *TRule) SetWriter(writer string) {
	t.writer = writer
}

func (t *TRule) SetDirection(direction string) error {
	if (direction != "LEFT") && (direction != "RIGHT") && (direction != "NOP") {
		return errors.New("Operation not recognized: " + direction)
	}
	t.direction = direction
	return nil
}
