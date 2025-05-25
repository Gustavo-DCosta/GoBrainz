package services

import (
	"fmt"

	"github.com/fatih/color"
)

func PromptOut() {
	var username string
	fmt.Print("Please insert an username: ")
	fmt.Scan(&username)

	color.RGB(227, 232, 104).Printf("GoBrainz@")
	fmt.Print(username)
	color.RGB(227, 232, 104).Print("$")

}
