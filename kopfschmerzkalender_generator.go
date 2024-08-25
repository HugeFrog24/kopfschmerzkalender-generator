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
	f.SetColWidth(sheetName, "A", "A", 5)
	f.SetColWidth(sheetName, "B", "K", 8)
	f.SetColWidth(sheetName, "L", "U", 6)
	f.SetColWidth(sheetName, "V", "X", 8)

	// Add title and logo
	f.SetCellValue(sheetName, "A1", "Kopfschmerzkalender")
	f.MergeCell(sheetName, "A1", "U1")
	titleStyle, _ := f.NewStyle(&excelize.Style{
		Font:      &excelize.Font{Bold: true, Size: 20, Italic: true},
		Alignment: &excelize.Alignment{Horizontal: "left", Vertical: "center"},
	})
	f.SetCellStyle(sheetName, "A1", "U1", titleStyle)

	f.SetCellValue(sheetName, "V1", "DMKG")
	f.MergeCell(sheetName, "V1", "X2")
	logoStyle, _ := f.NewStyle(&excelize.Style{
		Font:      &excelize.Font{Bold: true, Size: 16, Color: "FF0000"},
		Alignment: &excelize.Alignment{Horizontal: "center", Vertical: "center"},
	})
	f.SetCellStyle(sheetName, "V1", "X2", logoStyle)

	f.SetCellValue(sheetName, "V3", "Deutsche Migräne- und")
	f.SetCellValue(sheetName, "V4", "Kopfschmerzgesellschaft")
	f.SetCellValue(sheetName, "V5", "www.dmkg.de")
	f.MergeCell(sheetName, "V3", "X3")
	f.MergeCell(sheetName, "V4", "X4")
	f.MergeCell(sheetName, "V5", "X5")

	// Add medication, name, and month fields
	f.SetCellValue(sheetName, "A3", "Bitte vermerken Sie Ihre Medikamente,")
	f.SetCellValue(sheetName, "A4", "die Sie bei Kopfschmerzen einnehmen:")
	f.MergeCell(sheetName, "A3", "K3")
	f.MergeCell(sheetName, "A4", "K4")
	f.SetCellValue(sheetName, "A5", "A:")
	f.SetCellValue(sheetName, "A6", "B:")
	f.SetCellValue(sheetName, "A7", "C:")
	f.MergeCell(sheetName, "B5", "K5")
	f.MergeCell(sheetName, "B6", "K6")
	f.MergeCell(sheetName, "B7", "K7")

	// Set medication names from config
	f.SetCellValue(sheetName, "B5", config.MedicationA)
	f.SetCellValue(sheetName, "B6", config.MedicationB)
	f.SetCellValue(sheetName, "B7", config.MedicationC)

	f.SetCellValue(sheetName, "L3", "Name")
	f.MergeCell(sheetName, "M3", "U3")
	f.SetCellValue(sheetName, "M3", config.Name) // Set the name from config
	f.SetCellValue(sheetName, "L5", "Monat")
	f.MergeCell(sheetName, "M5", "U5")

	// Add main table headers
	headers := []string{"Tag", "Aus-löser", "Stärke", "Dauer (h)", "Pulsierend/ stechend", "Dumpf/ drückend", "Einseitig", "Beidseitig", "Vor-boten", "Erbrechen", "Übelkeit"}
	for i, header := range headers {
		cell := fmt.Sprintf("%c10", 'A'+i)
		f.SetCellValue(sheetName, cell, header)
	}

	// Add accompanying symptoms headers
	symptoms := []string{"Lärm-empfindl.", "Licht-empfindl.", "Geruchs-empfindl.", "Andere Symptome", "Medikament", "Tropfen/ Tabletten/ Zäpfchen", "Ja", "Nein", "Wenig"}
	for i, symptom := range symptoms {
		cell := fmt.Sprintf("%c10", 'L'+i)
		f.SetCellValue(sheetName, cell, symptom)
	}

	// Add new Oberbegriffe
	oberBegriffe := []string{"Schmerzart und Ort", "Begleitsymptome", "Anzahl der", "Hat Ihnen das Mittel geholfen?"}
	oberBegriffeColumns := []string{"A9:H9", "I9:O9", "P9:Q9", "R9:T9"}
	for i, begriff := range oberBegriffe {
		f.SetCellValue(sheetName, oberBegriffeColumns[i][:2], begriff)
		f.MergeCell(sheetName, oberBegriffeColumns[i][:2], oberBegriffeColumns[i][3:])
	}

	// Set style for Oberbegriffe
	oberBegriffeStyle, _ := f.NewStyle(&excelize.Style{
		Font:      &excelize.Font{Bold: true, Size: 10},
		Alignment: &excelize.Alignment{Horizontal: "center", Vertical: "center"},
		Fill:      excelize.Fill{Type: "pattern", Color: []string{"D9D9D9"}, Pattern: 1},
		Border:    []excelize.Border{{Type: "all", Color: "000000", Style: 1}},
	})
	f.SetCellStyle(sheetName, "A9", "T9", oberBegriffeStyle)

	// Add main table headers (combined with accompanying symptoms)
	headers = []string{
		"Tag", "Aus-löser", "Stärke", "Dauer (h)", "Pulsierend/ stechend", "Dumpf/ drückend", "Einseitig", "Beidseitig", "Vor-boten", "Erbrechen", "Übelkeit",
		"Lärm-empfindl.", "Licht-empfindl.", "Geruchs-empfindl.", "Andere Symptome", "Medikament", "Tropfen/ Tabletten/ Zäpfchen", "Ja", "Nein", "Wenig",
	}
	for i, header := range headers {
		cell := fmt.Sprintf("%c10", 'A'+i)
		f.SetCellValue(sheetName, cell, header)
	}

	// Add day numbers and create table grid
	for i := 1; i <= 31; i++ {
		row := 11 + i
		f.SetCellValue(sheetName, fmt.Sprintf("A%d", row), i)
		f.SetCellValue(sheetName, fmt.Sprintf("U%d", row), i)
		for col := 'A'; col <= 'T'; col++ {
			f.SetCellStyle(sheetName, fmt.Sprintf("%c%d", col, row), fmt.Sprintf("%c%d", col, row), getBorderedCellStyle(f))
		}
	}

	// Add legend
	f.SetCellValue(sheetName, "A43", "Schmerzstärke: 0-10 Punkte")
	f.SetCellValue(sheetName, "A44", "(0= kein Schmerz, 10= stärkster Schmerz)")

	f.SetCellValue(sheetName, "A46", "Vorboten:")
	f.SetCellValue(sheetName, "A47", "F  Flimmersehen")
	f.SetCellValue(sheetName, "A48", "G  Gefühlsstörung (Kribbeln, Pelzigkeit)")
	f.SetCellValue(sheetName, "A49", "S  Sprachstörung")
	f.SetCellValue(sheetName, "A50", "O Anderes Symptom:")

	f.SetCellValue(sheetName, "A52", "Dauer der Schmerzen:")
	f.SetCellValue(sheetName, "A53", "Geben Sie die Dauer in Stunden an")

	f.SetCellValue(sheetName, "A55", "Auslöser für Ihren Schmerz")
	f.SetCellValue(sheetName, "A56", "1. Aufregung /Stress")
	f.SetCellValue(sheetName, "A57", "2. Erholungsphase")
	f.SetCellValue(sheetName, "A58", "3.Änderung im Schlaf-Wach Rhythmus")
	f.SetCellValue(sheetName, "A59", "4. Menstruation")
	f.SetCellValue(sheetName, "A60", "5. Ihr persönlicher Auslöser")
	f.SetCellValue(sheetName, "A62", "6. Ein weiterer persönlicher Auslöser")

	f.SetCellValue(sheetName, "A64", "Andere Begleitsymptome:")
	f.SetCellValue(sheetName, "A65", "T  Augentränen")
	f.SetCellValue(sheetName, "A66", "R  Augenrötung")
	f.SetCellValue(sheetName, "A67", "N  Nasenlaufen / -Verstopfung")

	// Set month if provided
	if month != "" {
		f.SetCellValue(sheetName, "M5", month)
	}

	// Check config to add sample data
	if config.SampleData {
		addSampleData(f, sheetName, config)
	}

	// Add borders and styling to the entire table
	f.SetCellStyle(sheetName, "A9", "U10", getHeaderStyle(f))
	f.SetCellStyle(sheetName, "A11", "X42", getBorderedCellStyle(f))
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

		// Set "24 h" for duration
		f.SetCellValue(sheetName, fmt.Sprintf("D%d", row), "24 h")

		// Generate random strength between min and max intensity
		strength := r.Intn(config.MaxIntensity-config.MinIntensity+1) + config.MinIntensity
		f.SetCellValue(sheetName, fmt.Sprintf("C%d", row), strength)

		// Set "x" for Dumpf/drückend
		f.SetCellValue(sheetName, fmt.Sprintf("F%d", row), "x")

		// Set "x" for Lärm-empfindl., Licht-empfindl., and Geruchs-empfindl.
		f.SetCellValue(sheetName, fmt.Sprintf("L%d", row), "x")
		f.SetCellValue(sheetName, fmt.Sprintf("M%d", row), "x")
		f.SetCellValue(sheetName, fmt.Sprintf("N%d", row), "x")

		// Add medication data every minDays to maxDays
		if i == nextMedicationDay && len(validMedications) > 0 {
			// Set medication letter
			medication := validMedications[medicationIndex]
			f.SetCellValue(sheetName, fmt.Sprintf("P%d", row), medication)

			// Set helpfulness (randomly NOT HELPS or HELPS A BIT)
			if r.Intn(2) == 0 {
				f.SetCellValue(sheetName, fmt.Sprintf("S%d", row), "x") // NOT HELPS
			} else {
				f.SetCellValue(sheetName, fmt.Sprintf("T%d", row), "x") // HELPS A BIT
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
		return config, err
	}
	err = json.Unmarshal(data, &config)
	return config, err
}
