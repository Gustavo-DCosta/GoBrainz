package services

import (
	"fmt"

	"github.com/fatih/color"
)

const (
	StatementProfile = "GoBrainz profile"
	StatementAdd     = "GoBrainz add"
)

func PromptOut(username string) {
	color.RGB(227, 232, 104).Printf("GoBrainz@")
	fmt.Print(username)
	color.RGB(227, 232, 104).Print(" $ ")
}

func TokenCom(command string) {

	switch command {
	case StatementAdd:
		fmt.Println("ADDED page")
		break
	case StatementProfile:
		fmt.Println("This is your profile")
		break
	}
}
