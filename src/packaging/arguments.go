package packaging

// Arguments instruct
type Arguments struct {
	Install  []string `json:"install"`
	Remove   []string `json:"remove"`
	Shortcut []string `json:"shortcut"`
}

func NoArguments() *Arguments {
	return &Arguments{
		Install:  []string{},
		Remove:   []string{},
		Shortcut: []string{},
	}
}
