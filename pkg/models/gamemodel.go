package models

type OptionReqs struct {
	Mood      string
	Inventory []string
}

type GameOptions struct {
	ID          string
	Text        string
	Command     string
	Mood        string
	HasReqs     bool
	Requires    OptionReqs
	Inventory   string
	NextNode    string
	PlayType    string
	DeathNode   bool
	AfterAction bool
	StoryArc    string
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

// Default Game Node if a good Node ID cannot be returned from handler

var EmptyGameNode = GameNode{
	ID: "7331",
	Text: `
	--------------------------------
	EMBARRASSING ERROR # 38911C64-82
	--------------------------------

	ID: UNKNOWN
	ARC PHASE: FAILED...
	NODE SYNTHESIS: FAILED...
	ATTEMPTING BYPASS: INCOMPLETE

	PLAYER FAULT: FALSE
	DEVELOPER FAULT: TRUE
	
	STORY INCOMPLETE: TRUE
	GAME STATE RETRIEVAL: FAILED

	APOLOGIES DEAR PLAYER - YOUR SELECTED GAME PATH
	IS REGRETTABLY INCOMPLETE OR HAS FAILED TRANSFER
	PROTOCOLS. YOU CAN RESET THE GAME AND TRY A 
	DIFFERENT PATH FOR THE TIME BEING. 

	IF THIS BEHAVIOUR IS UNEXPECTED - PLEASE ALERT 
	THE OTTERLY TEAM VIA THE CONTACT PAGE...
	6 MILLION CREDITS HAVE BEEN DEPOSITED INTO 
	YOUR OTTERLY ACCOUNT* FOR THE INCONVENIENCE.

	*PLEASE NOTE - OTTERLY ACCOUNTS ARE ENTIRELY 
	FAKE

	--------------------------------
	END OF LINE
	--------------------------------
`,
	CodeNode:           false,
	CodeLength:         0,
	CodeFailedNextNode: 0,
	Options: []GameOptions{
		{
			ID:        "7331a",
			Text:      `Type "s" to Start Over`,
			Command:   "s",
			Mood:      "",
			HasReqs:   false,
			Inventory: "",
			NextNode:  "1",
		},
	},
	EarnedPoints: 0,
}
