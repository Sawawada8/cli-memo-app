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

// var filesPath string = os.Getenv("HOME") + "/.cli_memo_app/memos/"
var filesPath string = os.Getenv("HOME") + "/.cli_memo_app/demos/"

func Run() {
	fzfOption := NewOption().List

	switch len(os.Args) {
	case 1:
		// no args
		fmt.Println(stdOut(callFZF(fzfOption)))
		return
	case 3:
	case 4:
		// --height XX% -v or --height -v
		if IsContains(os.Args, []string{"--height", "-v"}) {
			fzfOption = append(fzfOption,
				"--preview",
				"cat "+filesPath+"{}",
			)
			if len(os.Args) == 4 {
				// args 4 の場合は、高さの指定有り
				fzfOption = AddHeight(fzfOption, os.Args[2])
			} else {
				fzfOption = AddHeight(fzfOption, "40%")
			}
			fmt.Println(stdOut(callFZF(fzfOption)))
			return
		}
	}

	// one args
	switch os.Args[1] {
	case "-v", "--view":
		fzfOption = append(fzfOption,
			"--preview",
			"cat "+filesPath+"{}",
		)
		fmt.Println(stdOut(callFZF(fzfOption)))
	case "-c", "--create":
		createMemo(new())
	case "-h", "--help":
		showHelp()
	case "--height":
		if len(os.Args) == 3 {
			fzfOption = AddHeight(fzfOption, os.Args[2])
		} else {
			fzfOption = AddHeight(fzfOption, "40%")
		}
		fmt.Println(stdOut(callFZF(fzfOption)))
	default:
		fmt.Println(os.Args[1], "は存在しないオプションです。\n-h, --help でオプション一覧を表示します。")
	}
}

func callFZF(fzfOption []string) string {
	out, err := pipeline.Output(
		[]string{"ls", filesPath},
		fzfOption,
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
	fmt.Println("オプション無し   : 既存のメモ一覧を表示します。")
	fmt.Println("-c, --create     : memoを作成します。")
	fmt.Println("-v, --view       : コンテンツを表示します。")
	fmt.Println("--height <xx%>   : FZF window の高さを指定します。(default: 40%)")
	fmt.Println("-h, --help       : helpを表示します。")
}

func new() map[string]string {
	scanner := bufio.NewScanner(os.Stdin)

	contents := map[string]string{
		"title": "",
		"body":  "",
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
