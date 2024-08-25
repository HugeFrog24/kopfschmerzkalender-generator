package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

func runGUI() {
	a := app.NewWithID("com.tibik.kopfschmerzkalender")
	w := a.NewWindow("Kopfschmerzkalender-Generator")

	// Create input fields
	minIntensityEntry := widget.NewEntry()
	maxIntensityEntry := widget.NewEntry()
	minDaysBetweenMedicationEntry := widget.NewEntry()
	maxDaysBetweenMedicationEntry := widget.NewEntry()
	nameEntry := widget.NewEntry()
	medicationAEntry := widget.NewEntry()
	medicationBEntry := widget.NewEntry()
	medicationCEntry := widget.NewEntry()
	outputFilePathEntry := widget.NewEntry()
	outputFilePathEntry.SetPlaceHolder("Leave empty for default path")

	// Create Browse button
	browseButton := widget.NewButton("Browse", func() {
		fd := dialog.NewFileSave(func(writer fyne.URIWriteCloser, err error) {
			if err != nil {
				dialog.ShowError(err, w)
				return
			}
			if writer == nil {
				return
			}
			defer writer.Close()
			outputFilePathEntry.SetText(writer.URI().Path())
		}, w)
		fd.SetFileName("Kopfschmerzkalender.xlsx")
		fd.SetFilter(storage.NewExtensionFileFilter([]string{".xlsx"}))
		fd.Show()
	})

	// Create output file path container
	outputFilePathContainer := container.NewBorder(nil, nil, nil, browseButton, outputFilePathEntry)

	sampleDataCheck := widget.NewCheck("Sample Data", func(checked bool) {
		// Toggle entry fields based on the checkbox state
		setEntryFieldsEnabled(checked, minIntensityEntry, maxIntensityEntry, minDaysBetweenMedicationEntry, maxDaysBetweenMedicationEntry, nameEntry, medicationAEntry, medicationBEntry, medicationCEntry)
	})

	// Set default values
	sampleDataCheck.SetChecked(true)
	// Enable entry fields initially since sample data is checked by default
	setEntryFieldsEnabled(true, minIntensityEntry, maxIntensityEntry, minDaysBetweenMedicationEntry, maxDaysBetweenMedicationEntry, nameEntry, medicationAEntry, medicationBEntry, medicationCEntry)
	minIntensityEntry.SetText("5")  // Pre-populate with 5
	maxIntensityEntry.SetText("10") // Pre-populate with 10
	minDaysBetweenMedicationEntry.SetText("5")
	maxDaysBetweenMedicationEntry.SetText("9")
	nameEntry.SetText("Max Mustermann")
	medicationAEntry.SetText("Ibuprofen 800")
	medicationBEntry.SetText("Thomapyrin Duo")

	// Create a slice to hold the selected months
	selectedMonths := make([]string, 0)

	// Create the months selection grid
	monthsGrid := createMonthsSelection(&selectedMonths)

	// Create form
	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Sample Data", Widget: sampleDataCheck},
			{Text: "Min Intensity (0-10)", Widget: minIntensityEntry},
			{Text: "Max Intensity (0-10)", Widget: maxIntensityEntry},
			{Text: "Min Days Between Medication", Widget: minDaysBetweenMedicationEntry},
			{Text: "Max Days Between Medication", Widget: maxDaysBetweenMedicationEntry},
			{Text: "Months", Widget: monthsGrid},
			{Text: "Name", Widget: nameEntry},
			{Text: "Medication A", Widget: medicationAEntry},
			{Text: "Medication B", Widget: medicationBEntry},
			{Text: "Medication C", Widget: medicationCEntry},
			{Text: "Output File Path", Widget: outputFilePathContainer},
		},
	}

	// Create start button
	startButton := widget.NewButton("Start", func() {
		// Validate intensity values
		minIntensity, minErr := validateIntensity(minIntensityEntry.Text)
		maxIntensity, maxErr := validateIntensity(maxIntensityEntry.Text)

		if minErr != nil || maxErr != nil {
			dialog.ShowError(fmt.Errorf("intensity must be a number between 0 and 10"), w)
			return
		}

		if minIntensity > maxIntensity {
			dialog.ShowError(fmt.Errorf("min intensity cannot be greater than max intensity"), w)
			return
		}

		// Validate medication delta time
		minDaysBetweenMedication, minDaysErr := parseInt(minDaysBetweenMedicationEntry.Text)
		maxDaysBetweenMedication, maxDaysErr := parseInt(maxDaysBetweenMedicationEntry.Text)

		if minDaysErr != nil || maxDaysErr != nil {
			dialog.ShowError(fmt.Errorf("days between medication must be valid numbers"), w)
			return
		}

		if minDaysBetweenMedication > maxDaysBetweenMedication {
			dialog.ShowError(fmt.Errorf("min days between medication cannot be greater than max days"), w)
			return
		}

		fmt.Printf("Selected months before creating config: %v\n", selectedMonths)

		config := Config{
			SampleData:               sampleDataCheck.Checked,
			MinDaysBetweenMedication: minDaysBetweenMedication,
			MaxDaysBetweenMedication: maxDaysBetweenMedication,
			Months:                   make([]string, len(selectedMonths)),
			Name:                     nameEntry.Text,
			MedicationA:              medicationAEntry.Text,
			MedicationB:              medicationBEntry.Text,
			MedicationC:              medicationCEntry.Text,
			OutputFilePath:           outputFilePathEntry.Text,
			MinIntensity:             minIntensity,
			MaxIntensity:             maxIntensity,
		}
		copy(config.Months, selectedMonths)

		fmt.Printf("Config months after creation: %v\n", config.Months)

		// Save config to file
		saveConfig(config)

		// Run the main program
		filePath, err := GenerateKopfschmerzkalender(config)
		if err != nil {
			dialog.ShowError(err, w)
		} else {
			showSuccessDialog(w, filePath)
		}
	})

	// Create exit button
	exitButton := widget.NewButton("Exit", func() {
		a.Quit()
	})

	// Create button container
	buttons := container.NewHBox(startButton, exitButton)

	// Create main container
	content := container.NewVBox(form, buttons)

	w.SetContent(content)
	w.Resize(fyne.NewSize(400, 400))
	w.ShowAndRun()
}

func parseInt(s string) (int, error) {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return i, nil
}

func saveConfig(config Config) {
	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	err = os.WriteFile("config.json", data, 0644)
	if err != nil {
		fmt.Println("Error writing config file:", err)
		return
	}

	fmt.Printf("Config saved successfully. Selected months: %v\n", config.Months)
}

func showSuccessDialog(w fyne.Window, filePath string) {
	openButton := widget.NewButton("Open File", func() {
		openFile(filePath)
	})

	content := container.NewVBox(
		widget.NewLabel("Kopfschmerzkalender generated successfully"),
		widget.NewLabel(fmt.Sprintf("File saved at: %s", filePath)),
		container.NewCenter(openButton),
	)

	dialog.ShowCustom("Success", "Close", content, w)
}

func openFile(path string) {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "start", "", path)
	case "darwin":
		cmd = exec.Command("open", path)
	default: // Linux and other Unix-like systems
		cmd = exec.Command("xdg-open", path)
	}
	err := cmd.Start()
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
	}
}

func setEntryFieldsEnabled(enabled bool, entries ...*widget.Entry) {
	for _, entry := range entries {
		if enabled {
			entry.Enable()
		} else {
			entry.Disable()
		}
		entry.Refresh()
	}
}

func validateIntensity(s string) (int, error) {
	i, err := strconv.Atoi(s)
	if err != nil || i < 0 || i > 10 {
		return 0, fmt.Errorf("intensity must be a number between 0 and 10")
	}
	return i, nil
}

// Helper function to create the months selection widget
func createMonthsSelection(selectedMonths *[]string) fyne.CanvasObject {
	months := []string{
		"Januar", "Februar", "März", "April", "Mai", "Juni",
		"Juli", "August", "September", "Oktober", "November", "Dezember",
	}

	checkGroup := widget.NewCheckGroup(months, func(selected []string) {
		*selectedMonths = selected
		fmt.Printf("Updated selected months: %v\n", *selectedMonths)
	})

	// Create a 3x4 grid layout
	grid := container.New(layout.NewGridLayout(3))

	for _, month := range months {
		check := widget.NewCheck(month, func(checked bool) {
			if checked {
				*selectedMonths = append(*selectedMonths, month)
			} else {
				for i, m := range *selectedMonths {
					if m == month {
						*selectedMonths = append((*selectedMonths)[:i], (*selectedMonths)[i+1:]...)
						break
					}
				}
			}
			checkGroup.SetSelected(*selectedMonths)
			fmt.Printf("Updated selected months: %v\n", *selectedMonths)
		})
		grid.Add(check)
	}

	// Set initial state based on config
	config, err := readConfig("config.json")
	if err == nil {
		checkGroup.SetSelected(config.Months)
		*selectedMonths = config.Months
	}

	return grid
}
