package services

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
)

func ParseConfigFile(path string, username string) (map[string]string, error) {
	file, err := os.Open(path)
	if err != nil {
		color.Yellow("File not found, creating new one: %s", path)
		file, err = os.Create(path)
		if err != nil {
			color.Red("Error creating config file: %v", err)
			return nil, err
		}
		defer file.Close()
		// Write username directly if file is new
		_, err = file.WriteString(fmt.Sprintf("username => %s\n", username))
		if err != nil {
			color.Red("Failed writing username to new file: %v", err)
			return nil, err
		}
		color.Green("Wrote default username to new config.")
		// Return immediately with only username
		return map[string]string{"username": username}, nil
	}
	defer file.Close()

	config := make(map[string]string)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		// skip empty lines, comment lines and version line
		if line == "" || strings.HasPrefix(line, "#") || strings.HasPrefix(line, "v") {
			continue
		}

		parts := strings.SplitN(line, "=>", 2)
		if len(parts) != 2 {
			color.Yellow("Invalid line: %s", line)
			continue
		}

		key := strings.TrimSpace(parts[0])
		fmt.Println("Debug: ", key)
		value := strings.TrimSpace(parts[1])
		fmt.Println("Debug: ", value)
		config[key] = value

	}
	return config, nil
}

func CheckForUsername(path string) {
	file, err := os.Open(path)
	if err != nil {
		color.Red("Can't open the file", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") || strings.HasPrefix(line, "v") {
			continue
		}
		username := strings.SplitN(line, "=>", 2)
		if username != nil {
			user := username[1]
			if user == " " {
				fmt.Println("This must be your first time, hello!")
			}
		}
	}
}

func WriteUsername(path string, Nusername string) {
	file, err := os.Open(path)
	if err != nil {
		color.Red("Can't open the file", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") || strings.HasPrefix(line, "v") {
			continue
		}
		username := strings.SplitN(line, "=>", 2)
		if username != nil || username[1] == " " {
			_, err := file.WriteString(fmt.Sprintf("username => %s\n", Nusername))
			if err != nil {
				color.Red("Error writting username", err)
			}
		} else {
			continue
		}
	}
}
