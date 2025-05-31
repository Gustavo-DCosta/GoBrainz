package main

import (
	"fmt"

	services "github.com/Gustavo-DCosta/GoBrainz/internal/Services"
	"github.com/fatih/color"
)

func main() {
	var username string
	color.Cyan(`        Welcome to GoBrainz
GoBrainz is a CLI tool to help manage Notion projects with simple commands`)

	color.Cyan(`Here are some basic commands to help you getting started with GoBrainz
	1. Help
	2. Exit`)

	fmt.Println("Please insert an username: ")
	fmt.Scan(&username)

	services.PromptOut(username)

}
