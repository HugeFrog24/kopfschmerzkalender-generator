package main

import (
	"log"
	"os"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "generate" {
		// Read config from file and generate Kopfschmerzkalender
		config, err := readConfig("config.json")
		if err != nil {
			log.Printf("Error reading config: %v", err)
			return
		}
		log.Printf("Read config: %+v", config)
		filePath, err := GenerateKopfschmerzkalender(config)
		if err != nil {
			log.Printf("Error generating Kopfschmerzkalender: %v", err)
			return
		}
		log.Printf("Kopfschmerzkalender generated successfully: %s", filePath)
	} else {
		// Run GUI
		runGUI()
	}
}
