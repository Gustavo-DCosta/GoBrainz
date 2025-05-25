package services

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
)

func EngineStart() {
	var username string

	fmt.Print("Please insert an username: ")
	fmt.Scan(&username)

	color.RGB(227, 232, 104).Printf("GoBrainz@")
	fmt.Print(username)
	color.RGB(227, 232, 104).Print("$")

	reader := bufio.NewReader(os.Stdin)
	cmdLine, _ := reader.ReadString('\n')
	cmdLine = strings.TrimSpace(cmdLine) // Remove that lil' newline

	tokens := strings.Split(cmdLine, " ")

	for i := 0; i < len(tokens); i++ {
		fmt.Println(tokens[i])
	}

}
