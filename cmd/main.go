package main

import (
	"fmt"
	"time"

	services "github.com/Gustavo-DCosta/GoBrainz/internal/Services"
	"github.com/fatih/color"
)

func main() {
	var name string
	var username string
	var choice int8
	color.Red("ENGINE STARTING...\n")
	time.Sleep(10 * time.Second)
	color.Cyan("Welcome to GoBrainz, a CLI assistant tool")

	color.Magenta("Before starting, would you like to select the syntax that you are most comfortable with?")
	color.Yellow(`
	1. @argument
	2. argument?
	3. GoBrainz argument
	`)

	fmt.Scan(&choice)

	switch choice {
	case 1:
		color.Green("Very well let's start...")
		services.CreatingCacheFile()
		break

	case 2:
		color.Green("Very well let's start...")
		services.CreatingCacheFile()
		break
	case 3:
		color.Green("Very well let's start...")
		services.CreatingCacheFile()
		break
	}

	color.RGB(227, 232, 104).Println("Please insert a name")
	fmt.Scan(&name)

	color.RGB(227, 232, 104).Println("Please insert a name")
	fmt.Scan(&username)

	services.PromptOut(name, username)
}
