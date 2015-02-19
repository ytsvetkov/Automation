package Automation

import "testing"

func TestCreationAndReading(t *testing.T) {
	rule1 := NewPRule("a", "1", "#", "a", "0")
	rule4 := NewPRule("a", "0", "#", "a", "1")
	rule2 := NewPRule("a", "1", "1!", "a", "0")
	rule6 := NewPRule("a", "0", "1!", "a", "1")
	rule3 := NewPRule("a", "1", "0!", "a", "0")
	rule5 := NewPRule("a", "0", "0!", "a", "1")

	book := NewEmptyDPRuleBook()
	book.AddRule(rule1)
	book.AddRule(rule2)
	book.AddRule(rule4)
	book.AddRule(rule6)
	book.AddRule(rule3)
	book.AddRule(rule5)

	states := SetFromSlice([]string{"a", "b"})
	accept := SetFromSlice([]string{"b"})

	stack := NewStack()
	stack.Push("#")

	dpda := NewDPDA("a", "b", "a", states, accept, stack, book)

	for _, char := range "000101010101010" {
		dpda.ReadCharacter(string(char))
	}

	theStack := dpda.GetStack()
	for _, letter := range "101010101010111" {
		bukva, ok := theStack.Peek()
		if !ok || bukva != string(letter) {
			t.Error("Bug-ish stack")
		}
		theStack.Pop()
	}
}
