package app

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/mattn/go-pipeline"
)

var filesPath string = os.Getenv("HOME") + "/.cli_memo_app/memos/"

func Run() {
	if len(os.Args) == 1 {
		// no args
		fmt.Println(stdOut(callFZF()))
		return
	}

	switch os.Args[1] {
	case "-c", "--create":
		createMemo(new())
	case "-h", "--help":
		showHelp()
	default:
		fmt.Println(os.Args[1], "は存在しないオプションです。")
	}
}

func callFZF() string {
	out, err := pipeline.Output(
		[]string{"ls", filesPath},
		[]string{
			"fzf",
			// "--height",
			// "40%",
			"--layout",
			"reverse",
			"--info",
			"inline",
			"--border",
			// "--preview",
			// "cat "+filesPath+"{}",
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

func createMemo(contents map[string]string) {
	fp, err := os.Create(filesPath + contents["title"] + ".txt")
	if err != nil {
		log.Fatal(err)
	}
	defer fp.Close()
	fp.WriteString(contents["body"])
}

func showHelp() {
	fmt.Println("cmemo <option>")
	fmt.Println("")
	fmt.Println("オプション無し : 既存のメモ一覧を表示します。")
	fmt.Println("-c, --create   : memoを作成します。")
	fmt.Println("-h, --help     : helpを表示します。")
}

func new() map[string]string {
	scanner := bufio.NewScanner(os.Stdin)

	contents := map[string]string{
		"title": "",
		"body": "",
	}

	for {
		if contents["title"] == "" {
			fmt.Print("title: ")
		} else {
			fmt.Print("body: ")
		}

		scanner.Scan()
		in := scanner.Text()

		if in == "/end" {
			break
		}

		if contents["title"] == "" {
			contents["title"] = in
		} else {
			contents["body"] += in + "\n"
			fmt.Println("title: " + contents["title"])
			fmt.Println("body: \n" + contents["body"])
			// return contents
		}
	}
	return contents
}
