package services

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
)

const (
	AddArgument = "add"
)

func EngineStart() {
	var username string

	fmt.Print("Please insert a username: ")
	fmt.Scan(&username)

	scanner := bufio.NewScanner(os.Stdin)

	for {
		color.RGB(227, 232, 104).Printf("GoBrainz@")
		fmt.Print(username)
		color.RGB(227, 232, 104).Print(" $ ")

		if !scanner.Scan() {
			break
		}

		input := scanner.Text()
		tokens := strings.Fields(input)

		// Compare against single-argument slice
		if len(tokens) == 1 && tokens[0] == AddArgument {
			color.Green("Notion page created")
		}
	}
}
