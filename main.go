package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "generate" {
		// Read config from file and generate Kopfschmerzkalender
		config, err := readConfig("config.json")
		if err != nil {
			fmt.Printf("Error reading config: %v\n", err)
			return
		}
		fmt.Printf("Read config: %+v\n", config)
		filePath, err := GenerateKopfschmerzkalender(config)
		if err != nil {
			fmt.Printf("Error generating Kopfschmerzkalender: %v\n", err)
			return
		}
		fmt.Printf("Kopfschmerzkalender generated successfully: %s\n", filePath)
	} else {
		// Run GUI
		runGUI()
	}
}
