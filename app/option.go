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
		// "--height",
		// "40%",
	}

	return &Option{
		List: list,
	}
}

func AddHeight(list []string, height string) []string {
	return append(list, "--height", height)
}