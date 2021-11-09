package main

import (
	"cliMemoApp/app"
	"fmt"
)

func main() {
	fmt.Println("\x1b[31m+---------------------------------------------------------+\x1b[0m")

	app.Run()

	fmt.Println("\x1b[31m+---------------------------------------------------------+\x1b[0m")
}
