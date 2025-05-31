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

		cmd := strings.Compare(token[1], StatementO)
		if cmd == -1 {
			color.Red("Error string does not contain GoBrainz keyword")
		}
	}
}
