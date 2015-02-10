package Automation

import "testing"

func TestPRuleFrom(t *testing.T) {
	rule := NewPRule("a", "1", "#", "a", "0")
	if "a" != rule.GetFrom() {
		t.Error("Problem with 'GetFrom' function for PRule class")
	}
	if "a" != rule.from {
		t.Error("Problem with initialising 'from' field for PRule class")
	}
	rule.SetFrom("from")
	if "from" != rule.GetFrom() {
		t.Error("Problem with 'SetFrom' function for PRule class")
	}
}

func TestPRuleWith(t *testing.T) {
	rule := NewPRule("a", "111", "#", "a", "0")
	if rule != nil {
		t.Error("Something went wrong with the initialisation!!!")
	}

	rule = NewPRule("a", "1", "#", "a", "0")
	if "1" != rule.GetWith() {
		t.Error("Problem with 'GetWith' function for PRule class")
	}

	if "1" != rule.with {
		t.Error("Problem with initialising 'with' field for PRule class")
	}

	err := rule.SetWith("00")
	if err == nil {
		t.Error("Something went wrong with the setting of 'With' !!!")
	}
	err = rule.SetWith("0")
	if err != nil {
		t.Error("Something went wrong with the setting of 'With' !!!")
	}
	if "0" != rule.GetWith() {
		t.Error("Problem with 'SetWith' function for PRule class")
		t.Error(rule)
	}
}

func TestPRuleTo(t *testing.T) {
	rule := NewPRule("a", "1", "#", "a", "0")

	if "a" != rule.GetTo() {
		t.Error("Problem with 'GetTo' function for PRule class")
	}

	if "a" != rule.to {
		t.Error("Problem with initialising 'to' field for PRule class")
	}

	rule.SetTo("from")
	if "from" != rule.GetTo() {
		t.Error("Problem with 'SetTo' function for PRule class")
	}
}

func TestPRulePush(t *testing.T) {
	rule := NewPRule("a", "1", "#", "a", "0")

	if "0" != rule.GetPush() {
		t.Error("Problem with 'GetPush' function for PRule class")
	}

	if "0" != rule.push {
		t.Error("Problem with initialising 'push' field for PRule class")
	}

	rule.SetPush("from")
	if "from" != rule.GetPush() {
		t.Error("Problem with 'SetPush' function for PRule class")
	}
}

func TestPRulePop(t *testing.T) {
	rule := NewPRule("a", "1", "#", "a", "0")

	if "#" != rule.GetPop() {
		t.Error("Problem with 'GetPop' function for PRule class")
	}

	if "#" != rule.pop {
		t.Error("Problem with initialising 'pop' field for PRule class")
	}

	rule.SetPop("from")
	if "from" != rule.GetPop() {
		t.Error("Problem with 'SetPop' function for PRule class")
	}
}
