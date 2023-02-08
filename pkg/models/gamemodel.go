package models

type OptionReqs struct {
	Mood      string
	Inventory []string
}

type GameOptions struct {
	ID        string
	Text      string
	Command   string
	Mood      string
	HasReqs   bool
	Requires  OptionReqs
	Inventory string
	NextNode  string
	PlayType  string
	DeathNode bool
}

type GameNode struct {
	ID                 string
	Text               string
	CodeNode           bool
	CodeLength         int
	CodeFailedNextNode int
	Options            []GameOptions
	EarnedPoints       int
}

type Game struct {
	ID          string
	Title       string
	TitleColor  string
	Image       string
	Description string
	IsAvailable bool
	Nodes       []GameNode
}
