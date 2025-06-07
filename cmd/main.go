package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	services "github.com/Gustavo-DCosta/GoBrainz/internal/Services"
	"github.com/fatih/color"
)

var username string

func init() {
	path := "./GoBrainz/Config.GoBrainz"

	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("⚠️ Config file not found. It will be created after username input.")
		return
	}

	username = strings.TrimSpace(string(data))
	fmt.Printf("Loaded username from config: %s\n", username)
}

func main() {
	var choice int8

	color.Cyan(`        Welcome to GoBrainz
		GoBrainz is a CLI tool to help manage Notion projects with simple commands`)

	color.Cyan(`Here are some basic commands to help you getting started with GoBrainz
			1. Help
			2. Exit`)

	fmt.Scan(&choice)
	switch choice {
	case 1:
		data, err := os.ReadFile("help.txt")
		if err != nil {
			color.Red("Error reading help.txt:", err)
			return
		}
		fmt.Print(string(data))
		fmt.Println("\n ")
	}

	// Only ask for username if it's not already set
	if username == "" {
		fmt.Print("Please insert a username: ")
		fmt.Scan(&username)
		services.WriteUsername(username)
	}

	for {
		services.PromptOut(username)

		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		services.TokenCom(input)
	}
}
