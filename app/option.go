package app

type Option struct {
	List []string;
}

func NewOption() *Option {
	// default argument
	list := []string{
		"fzf", 
		"--layout",
		"reverse",
		"--info",
		"inline",
		"--border",
	}

	return &Option{
		List: list,
	}
}
