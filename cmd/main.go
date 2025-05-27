package main

import (
	"errors"
	"fmt"
	"os"

	services "github.com/Gustavo-DCosta/GoBrainz/internal/Services"
	//"github.com/fatih/color"
)

func main() {
	_, err := os.Stat("syntax.txt")

	if errors.Is(err, os.ErrNotExist) {
		fmt.Println("Interal function working")
	}
	var choice int8

	services.PromptOut()

	fmt.Scan(&choice)

	switch choice {
	case 1:
		services.CreatingCacheFile(int(choice))
		services.EngineStart()
		break

	case 2:
		services.CreatingCacheFile(int(choice))
		services.EngineStart()
		break
	case 3:
		services.CreatingCacheFile(int(choice))
		services.EngineStart()
		break
	}
}
