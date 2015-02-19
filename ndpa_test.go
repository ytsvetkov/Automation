package Automation

import "testing"

func TestT(t *testing.T) {
	rule1 := NewPRule("a", "1", "#", "c", "0")
	rule7 := NewPRule("a", "1", "#", "c", "1")
	rule8 := NewPRule("a", "1", "#", "c", "2")
	rule9 := NewPRule("c", "1", "1", "a", "3")
	rule6 := NewPRule("c", "1", "2!", "a", "3")

	book := NewEmptyNPRuleBook()
	book.AddRule(rule1)
	book.AddRule(rule7)
	book.AddRule(rule8)
	book.AddRule(rule9)
	book.AddRule(rule6)

	states := SetFromSlice([]string{"a", "b", "c"})
	accept := SetFromSlice([]string{"b"})

	stack := NewStack()
	stack.PushString("#")

	a := NewSet()
	a.Add("a")
	b := NewSet()
	b.Add("b")
	c := NewSet()
	c.Add("c")

	npda := NewNPDA(a, b, a, states, accept, stack, book)

	npda.ReadCharacter("1")
	if npda.current.Intersection(c).Cardinality() != 1 {
		t.Error("Problem reading single charater!")
	}

	npda.Restart()
	npda.ReadString("11")

	if npda.current.Intersection(a).Cardinality() != 1 {
		t.Error("Problem reading multiple charaters!")
		t.Error(npda)
	}

	npda.ReadCharacter("8")
	if !npda.Rejecting() {
		t.Error("Not recognising rejection!")
	}
	if npda.Accepting() {
		t.Error("Not recognising rejection!")
	}
}
