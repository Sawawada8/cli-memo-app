package app

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/mattn/go-pipeline"
)


var filesPath string = os.Getenv("HOME")+"/.cli_memo_app/memos/"

func Run() {
	fmt.Println(stdOut(callFZF()))
}

func callFZF() string {
	out, err := pipeline.Output(
		[]string{"ls", filesPath},
		[]string{
            "fzf", 
            "--height",
            "40%",
            "--layout",
            "reverse",
            "--info",
            "inline",
            "--border",
            "--preview",
            "cat "+filesPath+"{}",
        },
	)
	if err != nil {
		log.Fatal(err)
	}
	return string(out)
}

func stdOut(fileName string) string {
	path := filesPath + fileName
	path = strings.TrimRight(path, "\n")

	cmd := exec.Command("cat", path)
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	return string(out)
}