package services

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

const (
	StatementO   = "GoBrainz"
	StatementADD = "Add"
)

func PromptOut(username string) {
	color.RGB(227, 232, 104).Printf("GoBrainz@")
	fmt.Print(username)
	color.RGB(227, 232, 104).Print(" $ ")
}

func TokenCom(command string) {
	token := strings.Split(command, " ")

	for i := 0; i < len(token); i++ {
		fmt.Println(token[i])
		if token[i] == StatementO+" " {
			color.Red("Syntax Error, must have the official statement 'GoBrainz' + 'argument' ")
		}
	}
}
