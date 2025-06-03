package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	services "github.com/Gustavo-DCosta/GoBrainz/internal/Services"
	"github.com/fatih/color"
	//	"github.com/gofiber/fiber/v2"
)

func init() {
	file, err := os.Create("Config.GoBrainz")

	if err != nil {
		color.Red("Error creating config file", err)
	}

	defer file.Close()
}

func main() {
	var username string
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

	fmt.Print("Please insert an username: ")
	fmt.Scan(&username)

	services.PromptOut(username)
	for {
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		services.TokenCom(input)
		services.PromptOut(username)
	}
}
