package services

import (
	"os"

	"github.com/fatih/color"
)

func WriteUsername(user string) {
	content := []byte(user)

	err := os.WriteFile("Config.GoBrainz", content, 0666)
	if err != nil {
		color.Red("Error writing username to file: %v", err)
	}
}

func ReadUsername() {

}
