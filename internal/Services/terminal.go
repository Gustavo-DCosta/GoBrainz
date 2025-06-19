package services

import (
	"fmt"
	"net/http"
	"os/exec"

	"github.com/fatih/color"
)

const (
	StatementProfile = "GoBrainz profile"
	StatementAdd     = "GoBrainz add"
	StatementLink    = "GoBrainz link"
)

func ServerOn() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from the server"))
	})

	err := http.ListenAndServe("127.0.0.1:8080", nil)
	if err != nil {
		color.RGB(239, 98, 110).Println(err)
	}
}

func PromptOut(username string) {
	color.RGB(227, 232, 104).Printf("GoBrainz@")
	fmt.Print(username)
	color.RGB(227, 232, 104).Print(" $ ")
}

func TokenCom(command string) {

	switch command {
	case StatementAdd:
		fmt.Println("ADDED page")
		cmd := exec.Command("node", "test1.js")
		output, err := cmd.CombinedOutput()

		if err != nil {
			color.RGB(239, 98, 110).Println("Failed creating a Notion block:\n", err)
			fmt.Print(string(output)) // Make sure to convert []byte to string
		}
		break
	case StatementProfile:
		fmt.Println("This is your profile")
		break
	case StatementLink:
		fmt.Println("Added link page")
		cmd := exec.Command("node", "linkblock.js")
		output, err := cmd.CombinedOutput()

		if err != nil {
			color.RGB(239, 98, 110).Println("Failed creating a Notion link block:\n", err)
			fmt.Print(string(output))
		}
		break
	}

}
