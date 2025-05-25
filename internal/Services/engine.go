package services

import (
	"github.com/fatih/color"
)

func EngineStart() {

	color.Cyan("Welcome to GoBrainz, a CLI assistant tool")

	color.Magenta("Before starting, would you like to select the syntax that you are most comfortable with?")
	color.Yellow(`
		1. @argument
		2. argument?
		3. GoBrainz argument
		`)

}
