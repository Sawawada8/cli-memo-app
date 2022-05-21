package main

import (
	"cliMemoApp/app"
	"fmt"
	"os/exec"
)

func main() {
	fmt.Println("\x1b[31m+---------------------------------------------------------+\x1b[0m")

	if existFZF() {
		app.Run()
	} else {
		fmt.Println("fzf が存在しません。")
		fmt.Println("ダウンロードして下さい。")
		fmt.Println("GitHubリポジトリ：https://github.com/junegunn/fzf")
	}

	fmt.Println("\x1b[31m+---------------------------------------------------------+\x1b[0m")
}

func existFZF() bool {
	output, _ := exec.Command("which", "fzf").Output()

	return len(output) != 0
}