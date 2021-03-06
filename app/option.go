package app

type Option struct {
	List []string
}

// defaultのfzfの設定オブジェクトを返す
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

// heightオプションを追加する
func AddHeight(list []string, height string) []string {
	return append(list, "--height", height)
}

func IsContains(args []string, words []string) bool {
	count := len(words)
	if count == 0 {
		return false
	}
	for i := 0; i < len(args); i++ {
		for ii := 0; ii < len(words); ii++ {
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
