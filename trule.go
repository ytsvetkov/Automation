package Automation

type TRule struct {
	with      string
	writer    string
	from      string
	to        string
	direction string
}

func NewTRule(from, with, to, writer, direction string) TRule {
	return TRule{from: from, with: with, to: to, writer: writer, direction: direction}
}

func NewTRules(from, with, to, writer, direction string) []TRule {
	rules := make([]TRule, 0, len(with))
	for _, char := range with {
		rules = append(rules, NewTRule(from, string(char), to, writer, direction))
	}
	return rules
}

func (t TRule) String() string {
	return "(" + t.from + ")-" + t.with + "-|(" + t.to + ")<" + t.writer + ">" + "[" + t.direction + "]"
}

func (t TRule) GetWith() string {
	return t.with
}

func (t TRule) GetWriter() string {
	return t.with
}

func (t TRule) GetTo() string {
	return t.with
}

func (t TRule) GetFrom() string {
	return t.with
}

func (t TRule) GetDirection() string {
	return t.with
}
