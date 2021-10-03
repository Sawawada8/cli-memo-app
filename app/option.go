package app

type Option struct {
	List []string;
}

func OptionNew() Option {
	// default argument
	list := []string{
		"fzf", 
		"--layout",
		"reverse",
		"--info",
		"inline",
		"--border",
	}

	return Option{
		List: list,
	}
}
