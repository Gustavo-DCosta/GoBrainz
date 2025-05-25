package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

func main() {
	var choice int8
	color.Cyan("Welcome to GoBrainz, a CLI assistant tool")
	
	color.Magenta("Before starting, would you like to select the syntax that you are most comfortable with?")
	color.Yellow(`
	1. @argument
	2. argument?
	3. GoBrainz argument
	`)

	fmt.Scan(&choice)

	for {
  	switch(choice) {
		case 1: 
			color.Green("Very well let's start...")
			file, err := os.Create("syntax.txt")

			if err != nil {
				fmt.Println("There was an error creating the cache file")
			}

			defer file.Close() 
			break
	
			case 2:
			color.Green("Very well let's start...")
			file, err := os.Create("syntax.txt")

			if err != nil {
				fmt.Println("Error creating cache file")
			}

			defer file.Close()
			break
	case 3:
			color.Green("Very well let's start...")
			file, err := os.Create("syntax.txt")

			if err != nil {
				fmt.Println("Error creating cache file")
			}
			defer file.Close()
			break
	}

	}

}
