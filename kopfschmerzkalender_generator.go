package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/xuri/excelize/v2"
)

func main() {
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	sheetName := "Kopfschmerzkalender"
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

	f.SetCellValue(sheetName, "L3", "Name")
	f.MergeCell(sheetName, "M3", "U3")
	f.SetCellValue(sheetName, "L5", "Monat")
	f.MergeCell(sheetName, "M5", "U5")

	// Add main table headers
	headers := []string{"Tag", "Aus-löser", "Stärke", "Dauer (h)", "Pulsierend/ stechend", "Dumpf/ drückend", "Einseitig", "Beidseitig", "Vor-boten", "Erbrechen", "Übelkeit"}
	for i, header := range headers {
		cell := fmt.Sprintf("%c9", 'A'+i)
		f.SetCellValue(sheetName, cell, header)
	}

	// Add accompanying symptoms headers
	symptoms := []string{"Lärm-empfindl.", "Licht-empfindl.", "Geruchs-empfindl.", "Andere Symptome", "Medikament", "Tropfen/ Tabletten/ Zäpfchen", "Ja", "Nein", "Wenig"}
	for i, symptom := range symptoms {
		cell := fmt.Sprintf("%c9", 'L'+i)
		f.SetCellValue(sheetName, cell, symptom)
	}

	// Add day numbers and create table grid
	for i := 1; i <= 31; i++ {
		row := 10 + i
		f.SetCellValue(sheetName, fmt.Sprintf("A%d", row), i)
		f.SetCellValue(sheetName, fmt.Sprintf("X%d", row), i)
		for col := 'A'; col <= 'X'; col++ {
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

	// Check for command-line argument to add sample data
	if len(os.Args) > 1 && os.Args[1] == "--sample" {
		addSampleData(f, sheetName)
		fmt.Println("Sample data added to the spreadsheet.")
	}

	// Add borders and styling to the entire table
	f.SetCellStyle(sheetName, "A9", "X9", getHeaderStyle(f))
	f.SetCellStyle(sheetName, "A10", "X41", getBorderedCellStyle(f))

	// Save the file
	if err := f.SaveAs("Kopfschmerzkalender.xlsx"); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Excel file created successfully: Kopfschmerzkalender.xlsx")
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

func addSampleData(f *excelize.File, sheetName string) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 1; i <= 31; i++ {
		row := 10 + i

		// Set "24h" for duration
		f.SetCellValue(sheetName, fmt.Sprintf("D%d", row), "24h")

		// Generate random strength between 5 and 10
		strength := r.Intn(6) + 5
		f.SetCellValue(sheetName, fmt.Sprintf("C%d", row), strength)

		// Set "x" for Dumpf/drückend
		f.SetCellValue(sheetName, fmt.Sprintf("F%d", row), "x")
	}
}
