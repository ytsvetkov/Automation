package Automation

import "testing"

func TestNewRule(t *testing.T) {
	rule, err := NewRule("r", "a", "q")
	if err != nil {
		t.Error("Something is wrong with rule")
		t.Error(err)
	} else if rule.GetStartingState() != "r" {
		t.Error("Expected starting state: r. Got: " + rule.GetStartingState())
	} else if rule.GetTransitionRule() != "a" {
		t.Error("Expected transition rule: a. Got: " + rule.GetTransitionRule())
	} else if rule.GetEndingState() != "q" {
		t.Error("Expected ending state: r. Got: " + rule.GetEndingState())
	}
}

func TestErrorNewRule(t *testing.T) {
	_, err := NewRule("p", "asd", "q")
	if err == nil {
		t.Error("Expected error. Got <nil>")
	}
}

func TestNewRules(t *testing.T) {
	str := "pqrst"
	rules := NewRules("a", str, "b")

	for i, rule := range rules {
		if rule.GetTransitionRule() != string(str[i]) {
			t.Error("There is a wrong rule! Expected: (a)-" + string(str[i]) + "-(b)" + " .Got: " + rule.String())
		}
	}
}
