package services

import (
	"fmt"
)

func PromptOut(name string, username string) {

	fmt.Print(name + "@" + username)
}
