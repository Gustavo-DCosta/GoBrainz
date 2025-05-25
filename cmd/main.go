package main

import (
	"fmt"

	services "github.com/Gustavo-DCosta/GoBrainz/internal/Services"
	//"github.com/fatih/color"
)

func main() {
	var choice int8

	services.PromptOut()

	fmt.Scan(&choice)

	switch choice {
	case 1:
		services.CreatingCacheFile()
		services.EngineStart()
		break

	case 2:
		services.CreatingCacheFile()
		services.EngineStart()
		break
	case 3:
		services.CreatingCacheFile()
		services.EngineStart()
		break
	}
}
