package packaging

// Arguments instruct
type Arguments struct {
	Install  []string `json:"install"`
	Remove   []string `json:"remove"`
	Run      []string `json:"run"`
	Shortcut []string `json:"shortcut"`
}

func NoArguments() *Arguments {
	return &Arguments{
		Install:  []string{},
		Remove:   []string{},
		Run:      []string{},
		Shortcut: []string{},
	}
}
