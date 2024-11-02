package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/xuri/excelize/v2"
)

func createSheet(f *excelize.File, sheetName, month string, config Config) {
	fmt.Printf("Creating sheet: %s, month: %s\n", sheetName, month)
	index, err := f.NewSheet(sheetName)
	if err != nil {
		fmt.Println(err)
		return
	}
	f.SetActiveSheet(index)

	// Set column widths
	if err := f.SetColWidth(sheetName, "A", "A", 5); err != nil {
		fmt.Printf("Error setting column width: %v\n", err)
	}
	if err := f.SetColWidth(sheetName, "B", "K", 8); err != nil {
		fmt.Printf("Error setting column width: %v\n", err)
	}
	if err := f.SetColWidth(sheetName, "L", "U", 6); err != nil {
		fmt.Printf("Error setting column width: %v\n", err)
	}
	if err := f.SetColWidth(sheetName, "V", "X", 8); err != nil {
		fmt.Printf("Error setting column width: %v\n", err)
	}

	// Add title and logo
	if err := f.SetCellValue(sheetName, "A1", "Kopfschmerzkalender"); err != nil {
		fmt.Printf("Error setting cell value: %v\n", err)
	}
	if err := f.MergeCell(sheetName, "A1", "U1"); err != nil {
		fmt.Printf("Error merging cells: %v\n", err)
	}
	titleStyle, _ := f.NewStyle(&excelize.Style{
		Font:      &excelize.Font{Bold: true, Size: 20, Italic: true},
		Alignment: &excelize.Alignment{Horizontal: "left", Vertical: "center"},
	})
	if err := f.SetCellStyle(sheetName, "A1", "U1", titleStyle); err != nil {
		fmt.Printf("Error setting cell style: %v\n", err)
	}

	if err := f.SetCellValue(sheetName, "V1", "DMKG"); err != nil {
		fmt.Printf("Error setting cell value: %v\n", err)
	}
	if err := f.MergeCell(sheetName, "V1", "X2"); err != nil {
		fmt.Printf("Error merging cells: %v\n", err)
	}
	logoStyle, _ := f.NewStyle(&excelize.Style{
		Font:      &excelize.Font{Bold: true, Size: 16, Color: "FF0000"},
		Alignment: &excelize.Alignment{Horizontal: "center", Vertical: "center"},
	})
	if err := f.SetCellStyle(sheetName, "V1", "X2", logoStyle); err != nil {
		fmt.Printf("Error setting cell style: %v\n", err)
	}

	if err := f.SetCellValue(sheetName, "V3", "Deutsche Migräne- und"); err != nil {
		fmt.Printf("Error setting cell value: %v\n", err)
	}
	if err := f.SetCellValue(sheetName, "V4", "Kopfschmerzgesellschaft"); err != nil {
		fmt.Printf("Error setting cell value: %v\n", err)
	}
	if err := f.SetCellValue(sheetName, "V5", "www.dmkg.de"); err != nil {
		fmt.Printf("Error setting cell value: %v\n", err)
	}
	if err := f.MergeCell(sheetName, "V3", "X3"); err != nil {
		fmt.Printf("Error merging cells: %v\n", err)
	}
	if err := f.MergeCell(sheetName, "V4", "X4"); err != nil {
		fmt.Printf("Error merging cells: %v\n", err)
	}
	if err := f.MergeCell(sheetName, "V5", "X5"); err != nil {
		fmt.Printf("Error merging cells: %v\n", err)
	}

	// Add medication, name, and month fields
	if err := f.SetCellValue(sheetName, "A3", "Bitte vermerken Sie Ihre Medikamente,"); err != nil {
		fmt.Printf("Error setting cell value: %v\n", err)
	}
	if err := f.SetCellValue(sheetName, "A4", "die Sie bei Kopfschmerzen einnehmen:"); err != nil {
		fmt.Printf("Error setting cell value: %v\n", err)
	}
	if err := f.MergeCell(sheetName, "A3", "K3"); err != nil {
		fmt.Printf("Error merging cells: %v\n", err)
	}
	if err := f.MergeCell(sheetName, "A4", "K4"); err != nil {
		fmt.Printf("Error merging cells: %v\n", err)
	}
	if err := f.SetCellValue(sheetName, "A5", "A:"); err != nil {
		fmt.Printf("Error setting cell value: %v\n", err)
	}
	if err := f.SetCellValue(sheetName, "A6", "B:"); err != nil {
		fmt.Printf("Error setting cell value: %v\n", err)
	}
	if err := f.SetCellValue(sheetName, "A7", "C:"); err != nil {
		fmt.Printf("Error setting cell value: %v\n", err)
	}
	if err := f.MergeCell(sheetName, "B5", "K5"); err != nil {
		fmt.Printf("Error merging cells: %v\n", err)
	}
	if err := f.MergeCell(sheetName, "B6", "K6"); err != nil {
		fmt.Printf("Error merging cells: %v\n", err)
	}
	if err := f.MergeCell(sheetName, "B7", "K7"); err != nil {
		fmt.Printf("Error merging cells: %v\n", err)
	}

	// Set medication names from config
	if err := f.SetCellValue(sheetName, "B5", config.MedicationA); err != nil {
		fmt.Printf("Error setting cell value: %v\n", err)
	}
	if err := f.SetCellValue(sheetName, "B6", config.MedicationB); err != nil {
		fmt.Printf("Error setting cell value: %v\n", err)
	}
	if err := f.SetCellValue(sheetName, "B7", config.MedicationC); err != nil {
		fmt.Printf("Error setting cell value: %v\n", err)
	}

	if err := f.SetCellValue(sheetName, "L3", "Name"); err != nil {
		fmt.Printf("Error setting cell value: %v\n", err)
	}
	if err := f.MergeCell(sheetName, "M3", "U3"); err != nil {
		fmt.Printf("Error merging cells: %v\n", err)
	}
	if err := f.SetCellValue(sheetName, "M3", config.Name); err != nil { // Set the name from config
		fmt.Printf("Error setting cell value: %v\n", err)
	}
	if err := f.SetCellValue(sheetName, "L5", "Monat"); err != nil {
		fmt.Printf("Error setting cell value: %v\n", err)
	}
	if err := f.MergeCell(sheetName, "M5", "U5"); err != nil {
		fmt.Printf("Error merging cells: %v\n", err)
	}

	// Add main table headers
	headers := []string{"Tag", "Aus-löser", "Stärke", "Dauer (h)", "Pulsierend/ stechend", "Dumpf/ drückend", "Einseitig", "Beidseitig", "Vor-boten", "Erbrechen", "Übelkeit"}
	for i, header := range headers {
		cell := fmt.Sprintf("%c10", 'A'+i)
		if err := f.SetCellValue(sheetName, cell, header); err != nil {
			fmt.Printf("Error setting cell value: %v\n", err)
		}
	}

	// Add accompanying symptoms headers
	symptoms := []string{"Lärm-empfindl.", "Licht-empfindl.", "Geruchs-empfindl.", "Andere Symptome", "Medikament", "Tropfen/ Tabletten/ Zäpfchen", "Ja", "Nein", "Wenig"}
	for i, symptom := range symptoms {
		cell := fmt.Sprintf("%c10", 'L'+i)
		if err := f.SetCellValue(sheetName, cell, symptom); err != nil {
			fmt.Printf("Error setting cell value: %v\n", err)
		}
	}

	// Add new Oberbegriffe
	oberBegriffe := []string{"Schmerzart und Ort", "Begleitsymptome", "Anzahl der", "Hat Ihnen das Mittel geholfen?"}
	oberBegriffeColumns := []string{"A9:H9", "I9:O9", "P9:Q9", "R9:T9"}
	for i, begriff := range oberBegriffe {
		if err := f.SetCellValue(sheetName, oberBegriffeColumns[i][:2], begriff); err != nil {
			fmt.Printf("Error setting cell value: %v\n", err)
		}
		if err := f.MergeCell(sheetName, oberBegriffeColumns[i][:2], oberBegriffeColumns[i][3:]); err != nil {
			fmt.Printf("Error merging cells: %v\n", err)
		}
	}

	// Set style for Oberbegriffe
	oberBegriffeStyle, _ := f.NewStyle(&excelize.Style{
		Font:      &excelize.Font{Bold: true, Size: 10},
		Alignment: &excelize.Alignment{Horizontal: "center", Vertical: "center"},
		Fill:      excelize.Fill{Type: "pattern", Color: []string{"D9D9D9"}, Pattern: 1},
		Border:    []excelize.Border{{Type: "all", Color: "000000", Style: 1}},
	})
	if err := f.SetCellStyle(sheetName, "A9", "T9", oberBegriffeStyle); err != nil {
		fmt.Printf("Error setting cell style: %v\n", err)
	}

	// Add main table headers (combined with accompanying symptoms)
	headers = []string{
		"Tag", "Aus-löser", "Stärke", "Dauer (h)", "Pulsierend/ stechend", "Dumpf/ drückend", "Einseitig", "Beidseitig", "Vor-boten", "Erbrechen", "Übelkeit",
		"Lärm-empfindl.", "Licht-empfindl.", "Geruchs-empfindl.", "Andere Symptome", "Medikament", "Tropfen/ Tabletten/ Zäpfchen", "Ja", "Nein", "Wenig",
	}
	for i, header := range headers {
		cell := fmt.Sprintf("%c10", 'A'+i)
		if err := f.SetCellValue(sheetName, cell, header); err != nil {
			fmt.Printf("Error setting cell value: %v\n", err)
		}
	}

	// Add day numbers and create table grid
	for i := 1; i <= 31; i++ {
		row := 11 + i
		if err := f.SetCellValue(sheetName, fmt.Sprintf("A%d", row), i); err != nil {
			fmt.Printf("Error setting cell value: %v\n", err)
		}
		if err := f.SetCellValue(sheetName, fmt.Sprintf("U%d", row), i); err != nil {
			fmt.Printf("Error setting cell value: %v\n", err)
		}
		for col := 'A'; col <= 'T'; col++ {
			if err := f.SetCellStyle(sheetName, fmt.Sprintf("%c%d", col, row), fmt.Sprintf("%c%d", col, row), getBorderedCellStyle(f)); err != nil {
				fmt.Printf("Error setting cell style: %v\n", err)
			}
		}
	}

	// Add legend
	if err := f.SetCellValue(sheetName, "A43", "Schmerzstärke: 0-10 Punkte"); err != nil {
		fmt.Printf("Error setting cell value: %v\n", err)
	}
	if err := f.SetCellValue(sheetName, "A44", "(0= kein Schmerz, 10= stärkster Schmerz)"); err != nil {
		fmt.Printf("Error setting cell value: %v\n", err)
	}

	if err := f.SetCellValue(sheetName, "A46", "Vorboten:"); err != nil {
		fmt.Printf("Error setting cell value: %v\n", err)
	}
	if err := f.SetCellValue(sheetName, "A47", "F  Flimmersehen"); err != nil {
		fmt.Printf("Error setting cell value: %v\n", err)
	}
	if err := f.SetCellValue(sheetName, "A48", "G  Gefühlsstörung (Kribbeln, Pelzigkeit)"); err != nil {
		fmt.Printf("Error setting cell value: %v\n", err)
	}
	if err := f.SetCellValue(sheetName, "A49", "S  Sprachstörung"); err != nil {
		fmt.Printf("Error setting cell value: %v\n", err)
	}
	if err := f.SetCellValue(sheetName, "A50", "O Anderes Symptom:"); err != nil {
		fmt.Printf("Error setting cell value: %v\n", err)
	}

	if err := f.SetCellValue(sheetName, "A52", "Dauer der Schmerzen:"); err != nil {
		fmt.Printf("Error setting cell value: %v\n", err)
	}
	if err := f.SetCellValue(sheetName, "A53", "Geben Sie die Dauer in Stunden an"); err != nil {
		fmt.Printf("Error setting cell value: %v\n", err)
	}

	if err := f.SetCellValue(sheetName, "A55", "Auslöser für Ihren Schmerz"); err != nil {
		fmt.Printf("Error setting cell value: %v\n", err)
	}
	if err := f.SetCellValue(sheetName, "A56", "1. Aufregung /Stress"); err != nil {
		fmt.Printf("Error setting cell value: %v\n", err)
	}
	if err := f.SetCellValue(sheetName, "A57", "2. Erholungsphase"); err != nil {
		fmt.Printf("Error setting cell value: %v\n", err)
	}
	if err := f.SetCellValue(sheetName, "A58", "3.Änderung im Schlaf-Wach Rhythmus"); err != nil {
		fmt.Printf("Error setting cell value: %v\n", err)
	}
	if err := f.SetCellValue(sheetName, "A59", "4. Menstruation"); err != nil {
		fmt.Printf("Error setting cell value: %v\n", err)
	}
	if err := f.SetCellValue(sheetName, "A60", "5. Ihr persönlicher Auslöser"); err != nil {
		fmt.Printf("Error setting cell value: %v\n", err)
	}
	if err := f.SetCellValue(sheetName, "A62", "6. Ein weiterer persönlicher Auslöser"); err != nil {
		fmt.Printf("Error setting cell value: %v\n", err)
	}

	if err := f.SetCellValue(sheetName, "A64", "Andere Begleitsymptome:"); err != nil {
		fmt.Printf("Error setting cell value: %v\n", err)
	}
	if err := f.SetCellValue(sheetName, "A65", "T  Augentränen"); err != nil {
		fmt.Printf("Error setting cell value: %v\n", err)
	}
	if err := f.SetCellValue(sheetName, "A66", "R  Augenrötung"); err != nil {
		fmt.Printf("Error setting cell value: %v\n", err)
	}
	if err := f.SetCellValue(sheetName, "A67", "N  Nasenlaufen / -Verstopfung"); err != nil {
		fmt.Printf("Error setting cell value: %v\n", err)
	}

	// Set month if provided
	if month != "" {
		if err := f.SetCellValue(sheetName, "M5", month); err != nil {
			fmt.Printf("Error setting cell value: %v\n", err)
		}
	}

	// Check config to add sample data
	if config.SampleData {
		addSampleData(f, sheetName, config)
	}

	// Add borders and styling to the entire table
	if err := f.SetCellStyle(sheetName, "A9", "U10", getHeaderStyle(f)); err != nil {
		fmt.Printf("Error setting cell style: %v\n", err)
	}
	if err := f.SetCellStyle(sheetName, "A11", "X42", getBorderedCellStyle(f)); err != nil {
		fmt.Printf("Error setting cell style: %v\n", err)
	}
}

func getHeaderStyle(f *excelize.File) int {
	style, _ := f.NewStyle(&excelize.Style{
		Font:      &excelize.Font{Bold: true, Size: 10},
		Alignment: &excelize.Alignment{Horizontal: "center", Vertical: "center", WrapText: true},
		Border:    []excelize.Border{{Type: "all", Color: "000000", Style: 1}},
		Fill:      excelize.Fill{Type: "pattern", Color: []string{"CCCCCC"}, Pattern: 1},
	})
	return style
}

func getBorderedCellStyle(f *excelize.File) int {
	style, _ := f.NewStyle(&excelize.Style{
		Border: []excelize.Border{{Type: "all", Color: "000000", Style: 1}},
	})
	return style
}

func addSampleData(f *excelize.File, sheetName string, config Config) {
	// #nosec G404 -- Using math/rand for non-cryptographic sample data generation
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	nextMedicationDay := 1
	medicationIndex := 0

	// Create a slice of valid medication letters and names
	var validMedications []string
	var medicationNames []string
	if config.MedicationA != "" {
		validMedications = append(validMedications, "A")
		medicationNames = append(medicationNames, config.MedicationA)
	}
	if config.MedicationB != "" {
		validMedications = append(validMedications, "B")
		medicationNames = append(medicationNames, config.MedicationB)
	}
	if config.MedicationC != "" {
		validMedications = append(validMedications, "C")
		medicationNames = append(medicationNames, config.MedicationC)
	}

	for i := 1; i <= 31; i++ {
		row := 11 + i

		// Set "Dauer (h)" to a random value between min and max
		duration := r.Intn(config.MaxDurationHours-config.MinDurationHours+1) + config.MinDurationHours
		if err := f.SetCellValue(sheetName, fmt.Sprintf("D%d", row), fmt.Sprintf("%d h", duration)); err != nil {
			fmt.Printf("Error setting cell value: %v\n", err)
		}

		// Generate random strength between min and max intensity
		strength := r.Intn(config.MaxIntensity-config.MinIntensity+1) + config.MinIntensity
		if err := f.SetCellValue(sheetName, fmt.Sprintf("C%d", row), strength); err != nil {
			fmt.Printf("Error setting cell value: %v\n", err)
		}

		// Set "Dumpf/drückend"
		if err := f.SetCellValue(sheetName, fmt.Sprintf("F%d", row), "x"); err != nil {
			fmt.Printf("Error setting cell value: %v\n", err)
		}

		// Set "Lärm-empfindl.", "Licht-empfindl.", "Geruchs-empfindl."
		if err := f.SetCellValue(sheetName, fmt.Sprintf("L%d", row), "x"); err != nil {
			fmt.Printf("Error setting cell value: %v\n", err)
		}
		if err := f.SetCellValue(sheetName, fmt.Sprintf("M%d", row), "x"); err != nil {
			fmt.Printf("Error setting cell value: %v\n", err)
		}
		if err := f.SetCellValue(sheetName, fmt.Sprintf("N%d", row), "x"); err != nil {
			fmt.Printf("Error setting cell value: %v\n", err)
		}

		// Add medication data every minDays to maxDays
		if i == nextMedicationDay && len(validMedications) > 0 {
			// Set medication letter
			medication := validMedications[medicationIndex]
			if err := f.SetCellValue(sheetName, fmt.Sprintf("P%d", row), medication); err != nil {
				fmt.Printf("Error setting cell value: %v\n", err)
			}

			// Set helpfulness (randomly NOT HELPS or HELPS A BIT)
			if r.Intn(2) == 0 {
				if err := f.SetCellValue(sheetName, fmt.Sprintf("S%d", row), "x"); err != nil { // NOT HELPS
					fmt.Printf("Error setting 'NOT HELPS' value: %v\n", err)
				}
			} else {
				if err := f.SetCellValue(sheetName, fmt.Sprintf("T%d", row), "x"); err != nil { // HELPS A BIT
					fmt.Printf("Error setting 'HELPS A BIT' value: %v\n", err)
				}
			}

			// Update next medication day and medication index
			nextMedicationDay += r.Intn(config.MaxDaysBetweenMedication-config.MinDaysBetweenMedication+1) + config.MinDaysBetweenMedication
			medicationIndex = (medicationIndex + 1) % len(validMedications)
		}
	}

	// Modify the log statement
	if len(validMedications) > 0 {
		fmt.Printf("Sample data added to the spreadsheet '%s'. Medication applied every %d to %d days. Medications: %s\n",
			sheetName, config.MinDaysBetweenMedication, config.MaxDaysBetweenMedication, strings.Join(medicationNames, ", "))
	} else {
		fmt.Printf("Sample data added to the spreadsheet '%s'. No medication applied.\n", sheetName)
	}
}

func readConfig(filename string) (Config, error) {
	var config Config
	data, err := os.ReadFile(filename)
	if err != nil {
		return config, fmt.Errorf("error reading config file: %w", err)
	}
	err = json.Unmarshal(data, &config)
	return config, err
}
