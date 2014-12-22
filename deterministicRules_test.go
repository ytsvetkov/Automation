package Automation

import "testing"

func TestNewBook(t *testing.T) {
	rules := NewRules("from", "disowdh", "to")
	_, err1 := NewDeterministicRuleBook(rules)

	if err1 != nil {
		t.Error("Unexpected error")
		t.Error(err1)
	}

	r1, _ := NewRule("a", "b", "c")
	r2, _ := NewRule("a", "d", "c")
	r3, _ := NewRule("a", "d", "b")
	_, err2 := NewDeterministicRuleBook([]*Rule{r1, r2, r3})
	if err2 == nil {
		t.Error("Expected non-deterministic behaviour detection. Got <nil>")
	}

}

func TestGetRule(t *testing.T) {
	rules := NewRules("from", "disowdh", "to")
	book1, _ := NewDeterministicRuleBook(rules)

	rule := book1.GetRule("from", "d")
	if rule == nil {
		t.Error("Rule not found.")
	}

	rule2 := book1.GetRule("from", "j")
	if rule2 != nil {
		t.Error("Found not existing rule!")
	}

	rule3, _ := NewRule("from", "j", "to")
	book1.AddRule(rule3)
	if rule4 := book1.GetRule("from", "j"); rule4 == nil {
		t.Error("Rule not found.")
	}
}

func TestAddRule(t *testing.T) {
	rules := NewRules("from", "disowdh", "to")
	book1, _ := NewDeterministicRuleBook(rules)

	rule, _ := NewRule("a", "b", "c")
	book1.AddRule(rule)
	if get := book1.GetRule("a", "b"); get == nil {
		t.Error("Existing rule not found!")
	} else if get.GetStartingState() != "a" || get.GetTransitionRule() != "b" || get.GetEndingState() != "c" {
		t.Error("Incorect rule found!")
	}

	if get := book1.GetRule("a", "c"); get != nil {
		t.Error("Non-existing rule found!")
	}
}
