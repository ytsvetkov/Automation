package Automation

import "testing"

func TestRuleFrom(t *testing.T) {
	rule := NewRule("mu", "s", "aka")
	if "mu" != rule.GetFrom() {
		t.Error("Problem with 'GetFrom' function for Rule class")
	}
	if "mu" != rule.from {
		t.Error("Problem with initialising 'from' field for Rule class")
	}
	rule.SetFrom("from")
	if "from" != rule.GetFrom() {
		t.Error("Problem with 'SetFrom' function for Rule class")
	}
}

func TestRuleWith(t *testing.T) {
	rule := NewRule("mu", "sa", "ka")
	if rule != nil {
		t.Error("Something went wrong with the initialisation!!!")
	}

	rule = NewRule("mu", "s", "aka")
	if "s" != rule.GetWith() {
		t.Error("Problem with 'GetWith' function for Rule class")
	}

	if "s" != rule.with {
		t.Error("Problem with initialising 'with' field for Rule class")
	}

	err := rule.SetWith("from")
	if err == nil {
		t.Error("Something went wrong with the setting of 'With' !!!")
	}
	err = rule.SetWith("f")
	if err != nil {
		t.Error("Something went wrong with the setting of 'With' !!!")
	}
	if "f" != rule.GetWith() {
		t.Error("Problem with 'SetWith' function for Rule class")
	}
}

func TestRuleTo(t *testing.T) {
	rule := NewRule("mu", "s", "aka")

	if "aka" != rule.GetTo() {
		t.Error("Problem with 'GetTo' function for Rule class")
	}

	if "aka" != rule.to {
		t.Error("Problem with initialising 'to' field for Rule class")
	}

	rule.SetTo("from")
	if "from" != rule.GetTo() {
		t.Error("Problem with 'SetTo' function for Rule class")
	}
}
