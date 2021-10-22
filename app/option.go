package app

type Option struct {
	List []string
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

func IsContains(args []string, words []string) bool {
	count := len(words)
	for i := 0; i < len(args); i++ {
		for ii := 0; ii < len(words); i++ {
			if args[i] == words[ii] {
				count--
				break
			}
		}
		if count <= 0 {
			return true
		}
	}
	return false
}
