package Automation

import "testing"

func TestNewNBook(t *testing.T) {
	rules := NewRules("from", "disowdh", "to")
	book1 := NewNonDeterministicRuleBook(rules)

	if len(book1["from"]) != 6 {
		t.Error("Missmatch of rules!")
	}

}

func TestAddNRule(t *testing.T) {
	rules := NewRules("from", "disowdh", "to")
	book1 := NewNonDeterministicRuleBook(rules)

	rule, _ := NewRule("a", "b", "c")
	book1.AddRule(rule)

	if r := book1.GetRulesEnd("a", "b"); r == nil {
		t.Error("Rule not found")
	}
}

func TestGetRules(t *testing.T) {
	rules := NewRules("from", "disowdh", "to")
	book1 := NewNonDeterministicRuleBook(rules)

	if x := book1.GetRulesEnd("from", "d"); len(x) != 1 {
		t.Error("Problem")
	}
}
