package Automation

type RuleBook interface {
	AddRule(r *Rule) error
	AddRules(r []*Rule) error
	GetFromState(from string) [][2]string
	GetRuleEnd(from, with string) Set
	GetFromTransition(from string) Set
	String() string
	GetAllRules() []*Rule
}

type RegularAutomata interface {
	ReadCharacter(char string)
	ReadString(word string)
	Accepting() bool
	// String() string
	GetAlphabet() Set
	GetAllStates() Set
	GetAllTransitionsFor(from, with string) Set
	GetFromState(from string) [][2]string
	GetAllEnds() Set
	GetStartStates() Set
	GetAcceptStates() Set
	GetReject() string
	GetAllRules() []*Rule
}
