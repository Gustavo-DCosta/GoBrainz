package services

import (
	"fmt"

	"github.com/fatih/color"
)

func PromptOut(username string) {
	color.RGB(227, 232, 104).Printf("GoBrainz@")
	fmt.Print(username)
	color.RGB(227, 232, 104).Print(" $ ")
}
