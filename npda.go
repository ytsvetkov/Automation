package Automation

import "fmt"
import "regexp"
import "strings"

type NPDA struct {
	start   Set
	reject  Set
	current Set
	states  Set
	accept  Set
	stack   *Stack
	rules   NPRuleBook
}

//Returns new Non-deterministic Pushdown Automata
func NewNPDA(start, reject, current, states, accept Set, stack *Stack, rules NPRuleBook) *NPDA {
	return &NPDA{start: start, reject: reject, current: current, states: states, accept: accept, stack: stack, rules: rules}
}

func (n *NPDA) Restart() {
	for i := 0; i < n.stack.Len(); i++ {
		_, _ = n.stack.Pop()
	}
	n.stack.Push("#")
	n.current = n.start
}

func (n *NPDA) String() string {
	return "Current state: " + n.current.String() + "\nState of the stack: " + n.stack.String()
}

func (n *NPDA) Accepting() bool {
	return !(n.accept.Intersection(n.current).Cardinality() != 0)
}

func (n *NPDA) Rejecting() bool {
	return n.reject.Intersection(n.current).Cardinality() != 0
}

func (n *NPDA) ReadCharacter(char string) {
	stackTop, err := n.stack.Peek()
	if !err {
		n.current = n.reject
		return
	}
	stackSlice := regexp.MustCompile(", ").Split(stackTop, -1)

	pop := map[string]int{}
	rule := [][][2]string{}
	rule_i := [][2]string{}

	for _, state := range n.current.Values() {
		for _, letter := range stackSlice {
			rule_i = n.rules.GetRuleEnd(state, char, letter)
			pop[letter] += 1
			if len(rule_i) == 0 {
				rule_i = n.rules.GetRuleEnd(state, char, letter+"!")
				pop[letter] -= 1
			}

			if len(rule_i) != 0 {
				rule = append(rule, rule_i)
			}

		}
	}

	if len(rule) == 0 {
		n.current = n.reject
		return
	}

	n.current = NewSet()
	for index, states := range rule {
		for index_i, _ := range states {
			// fmt.Println(rule)
			n.current.Add(rule[index][index_i][0])
		}
	}
	// fmt.Println(n.current)

	str, err := n.stack.Pop()
	if err != true {
		str = ""
	}

	for letter, count := range pop {
		if count > 0 && len(stackSlice) == 1 {
			str = strings.Replace(str, letter, "", count)
		} else if count > 0 && len(stackSlice) >= 1 {
			str = strings.Replace(str, letter+", ", "", count)
		}
	}

	if strings.Count(str, ",") >= 1 {
		str += ", "
	}
	for index, states := range rule {
		for index_i, _ := range states {
			if rule[index][index_i][1] != " " {
				str += rule[index][index_i][1] + ", "
			}
		}
	}

	if string(str[len(str)-1]) == " " {
		str = str[:len(str)-2]
	}

	n.stack.PushString(str)
}

func (n *NPDA) ReadString(word string) {
	for index, char := range word {
		if n.Rejecting() {
			fmt.Println(index)
			return
		}
		n.ReadCharacter(string(char))
	}
}

//Returns the string representation of stack.
//The format is '[' member1 ',' member2 ',' ... ']'
func (n *NPDA) GetStackAsString() string {
	return n.stack.String()
}

//Returns the stack of the machine.
func (n *NPDA) GetStack() Stack {
	return *(n.stack)
}
