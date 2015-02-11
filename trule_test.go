package Automation

import "testing"

func TestTRuleFrom(t *testing.T) {
	rule := NewTRule("alpha", "0", "beta", "147", "LEFT")
	if "alpha" != rule.GetFrom() {
		t.Error("Problem with 'GetFrom' function for TRule class")
	}
	if "alpha" != rule.from {
		t.Error("Problem with initialising 'from' field for TRule class")
	}
	rule.SetFrom("from")
	if "from" != rule.GetFrom() {
		t.Error("Problem with 'SetFrom' function for TRule class")
	}
}

func TestTRuleTo(t *testing.T) {
	rule := NewTRule("alpha", "0", "beta", "147", "LEFT")
	if "beta" != rule.GetTo() {
		t.Error("Problem with 'GetTo' function for TRule class")
	}
	if "beta" != rule.to {
		t.Error("Problem with initialising 'to' field for TRule class")
	}
	rule.SetTo("from")
	if "from" != rule.GetTo() {
		t.Error("Problem with 'SetTo' function for TRule class")
	}
}

func TestTRuleWriter(t *testing.T) {
	rule := NewTRule("alpha", "0", "beta", "147", "LEFT")
	if "alpha" != rule.GetFrom() {
		t.Error("Problem with 'GetFrom' function for TRule class")
	}
	if "alpha" != rule.from {
		t.Error("Problem with initialising 'from' field for TRule class")
	}
	rule.SetFrom("from")
	if "from" != rule.GetFrom() {
		t.Error("Problem with 'SetFrom' function for TRule class")
	}
}

func TestTRuleDirection(t *testing.T) {
	rule := NewTRule("alpha", "0", "beta", "147", "LEFT")
	if "LEFT" != rule.GetDirection() {
		t.Error("Problem with 'GetDirection' function for TRule class")
	}
	if "LEFT" != rule.direction {
		t.Error("Problem with initialising 'direction' field for TRule class")
	}
	rule.SetDirection("RIGHT")
	if "RIGHT" != rule.GetDirection() {
		t.Error("Problem with 'SetDirection' function for TRule class")
	}
}

func TestTRuleWith(t *testing.T) {
	rule := NewTRule("alpha", "000", "beta", "147", "LEFT")
	if rule != nil {
		t.Error("Problem with initialisation")
	}

	rule = NewTRule("alpha", "0", "beta", "147", "LEFT")
	if "0" != rule.GetWith() {
		t.Error("Problem with 'GetWith' function for TRule class")
	}
	if "0" != rule.with {
		t.Error("Problem with 'with' field for TRule class")
	}

	err := rule.SetWith("11")
	if err == nil {
		t.Error("Something went wrong with the setting of 'With' in TRule class!!!")
	}
	err = rule.SetWith("1")
	if err != nil {
		t.Error("Something went wrong with the setting of 'With' in TRule class!!!")
	}
	if "1" != rule.GetWith() {
		t.Error("Problem with 'SetWith' function for TRule class")
	}
}
