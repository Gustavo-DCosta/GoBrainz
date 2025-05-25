package services

import (
	"os"

	"github.com/fatih/color"
)

func CreatingCacheFile() { //fancy name
	file, err := os.Create("syntax.txt")

	if err != nil {
		color.Red("Error creating cache file: ", err)
	}

	defer file.Close()
}
