package services

import (
	"os"

	"github.com/fatih/color"
)

func CreatingCacheFile(choice int) { //fancy name

	switch choice {
	case 1:
		file, err := os.Create("syntax.txt")

		if err != nil {
			color.Red("Error creating cache syntax file", err)
		}

		err = os.WriteFile("syntax.txt", []byte("Syntax Flavour: @argument"), 0666)

		if err != nil {
			color.Red("Error writing configurations on syntax file", err)
		}

		defer file.Close()
	case 2:
		file, err := os.Create("syntax.txt")

		if err != nil {
			color.Red("Error creating cache syntax file", err)
		}

		err = os.WriteFile("syntax.txt", []byte("Syntax Flavour: argument?"), 0666)

		if err != nil {
			color.Red("Error writing configurations on syntax file", err)
		}

		defer file.Close()

	case 3:
		file, err := os.Create("syntax.txt")

		if err != nil {
			color.Red("Error creating cache syntax file", err)
		}

		err = os.WriteFile("syntax.txt", []byte("Syntax Flavour: GoBrainz argument"), 0666)

		if err != nil {
			color.Red("Error writing configurations on syntax file", err)
		}

		defer file.Close()
	}

}
