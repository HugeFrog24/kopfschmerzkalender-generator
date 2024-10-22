package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/xuri/excelize/v2"
)

const GithubRepoURL = "https://github.com/HugeFrog24/kopfschmerzkalender-generator"

type Config struct {
	SampleData               bool     `json:"sample_data"`
	MinDaysBetweenMedication int      `json:"min_days_between_medication"`
	MaxDaysBetweenMedication int      `json:"max_days_between_medication"`
	Months                   []string `json:"months"`
	Name                     string   `json:"name"`
	MedicationA              string   `json:"medication_a"`
	MedicationB              string   `json:"medication_b"`
	MedicationC              string   `json:"medication_c"`
	OutputFilePath           string   `json:"output_file_path"`
	MinIntensity             int      `json:"min_intensity"`
	MaxIntensity             int      `json:"max_intensity"`
	// New fields for duration hours
	MinDurationHours int `json:"min_duration_hours"`
	MaxDurationHours int `json:"max_duration_hours"`
}

// GenerateKopfschmerzkalender is declared here but implemented in kopfschmerzkalender_generator.go
func GenerateKopfschmerzkalender(config Config) (string, error) {
	fmt.Printf("Generating Kopfschmerzkalender with config: %+v\n", config)
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println("Error closing Excel file:", err)
		}
	}()

	// Remove the default Sheet1
	if err := f.DeleteSheet("Sheet1"); err != nil {
		fmt.Printf("Warning: Failed to delete default Sheet1: %v\n", err)
	}

	fmt.Printf("Generating sheets for months: %v\n", config.Months)

	// Remove the unnecessary nil check
	if len(config.Months) == 0 {
		createSheet(f, "Kopfschmerzkalender", "", config)
	} else {
		for _, month := range config.Months {
			createSheet(f, month, month, config)
		}
	}

	// Use the custom output file path if provided, otherwise use the default
	var filePath string
	if config.OutputFilePath != "" {
		filePath = config.OutputFilePath
	} else {
		// Get the current working directory
		currentDir, err := os.Getwd()
		if err != nil {
			return "", fmt.Errorf("error getting current directory: %v", err)
		}
		filePath = filepath.Join(currentDir, "Kopfschmerzkalender.xlsx")
	}

	// Save the file
	if err := f.SaveAs(filePath); err != nil {
		return "", fmt.Errorf("error saving Excel file: %v", err)
	}

	fmt.Printf("Excel file created successfully: %s\n", filePath)
	return filePath, nil
}
