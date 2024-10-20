package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"os"
	"os/exec"
	"runtime"
	"strconv"

	l "github.com/HugeFrog24/kopfschmerzkalender-generator/localization"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
	"golang.org/x/text/language"
)

// Helper function to create a range input container
func createRangeField(minEntry, maxEntry *widget.Entry) *fyne.Container {
	return container.NewHBox(
		minEntry,
		widget.NewLabel(" - "),
		maxEntry,
	)
}

// Add this function to create a language selector
func createLanguageSelector(updateUI func()) *widget.Select {
	languages := []string{"English", "Deutsch"}
	langSelect := widget.NewSelect(languages, func(selected string) {
		switch selected {
		case "English":
			l.SetLanguage(language.English)
			log.Println("Language changed to English") // Add logging for language change
		case "Deutsch":
			l.SetLanguage(language.German)
			log.Println("Language changed to German") // Add logging for language change
		}
		// Call the updateUI function to refresh all widgets
		updateUI()
	})
	langSelect.SetSelected("Deutsch") // Set default language
	return langSelect
}

func runGUI() {
	a := app.NewWithID("com.tibik.kopfschmerzkalender")
	w := a.NewWindow(l.T(l.MsgAppTitle))

	// Create input fields
	minIntensityEntry := widget.NewEntry()
	maxIntensityEntry := widget.NewEntry()
	minDaysBetweenMedicationEntry := widget.NewEntry()
	maxDaysBetweenMedicationEntry := widget.NewEntry()

	// New entries for duration hours
	minDurationHoursEntry := widget.NewEntry()
	maxDurationHoursEntry := widget.NewEntry()

	nameEntry := widget.NewEntry()
	nameEntry.SetPlaceHolder(l.T(l.MsgNamePlaceholder))
	medicationAEntry := widget.NewEntry()
	medicationAEntry.SetPlaceHolder(l.T(l.MsgMedicationAPlaceholder))
	medicationBEntry := widget.NewEntry()
	medicationBEntry.SetPlaceHolder(l.T(l.MsgMedicationBPlaceholder))
	medicationCEntry := widget.NewEntry()
	medicationCEntry.SetPlaceHolder(l.T(l.MsgMedicationCPlaceholder))
	outputFilePathEntry := widget.NewEntry()
	outputFilePathEntry.SetPlaceHolder(l.T(l.MsgOutputFilePathPlaceholder))

	// Create Browse button
	browseButton := widget.NewButton(l.T(l.MsgBrowse), func() {
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

	sampleDataCheck := widget.NewCheck(l.T(l.MsgSampleData), func(checked bool) {
		// Log the state change
		if checked {
			log.Println("Sample data activated")
		} else {
			log.Println("Sample data deactivated")
		}
		// Toggle entry fields based on the checkbox state
		setEntryFieldsEnabled(checked,
			minIntensityEntry, maxIntensityEntry,
			minDaysBetweenMedicationEntry, maxDaysBetweenMedicationEntry,
			minDurationHoursEntry, maxDurationHoursEntry, // Include new fields
			nameEntry, medicationAEntry, medicationBEntry, medicationCEntry)
	})

	// Set default values
	sampleDataCheck.SetChecked(true)
	// Enable entry fields initially since sample data is checked by default
	setEntryFieldsEnabled(true,
		minIntensityEntry, maxIntensityEntry,
		minDaysBetweenMedicationEntry, maxDaysBetweenMedicationEntry,
		minDurationHoursEntry, maxDurationHoursEntry, // Include new fields
		nameEntry, medicationAEntry, medicationBEntry, medicationCEntry)
	minIntensityEntry.SetText("5")  // Pre-populate with 5
	maxIntensityEntry.SetText("10") // Pre-populate with 10
	minDaysBetweenMedicationEntry.SetText("5")
	maxDaysBetweenMedicationEntry.SetText("9")
	minDurationHoursEntry.SetText("1")  // Example default min hours
	maxDurationHoursEntry.SetText("24") // Example default max hours

	// Create a slice to hold the selected months
	selectedMonths := make([]string, 0)

	// Create the months selection grid
	monthsGrid := createMonthsSelection(&selectedMonths)

	// Group range fields into horizontal containers
	intensityRange := createRangeField(minIntensityEntry, maxIntensityEntry)
	daysBetweenMedRange := createRangeField(minDaysBetweenMedicationEntry, maxDaysBetweenMedicationEntry)
	durationHoursRange := createRangeField(minDurationHoursEntry, maxDurationHoursEntry) // New range field

	// Create form
	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: l.T(l.MsgSampleData), Widget: sampleDataCheck},
			{Text: l.T(l.MsgIntensity), Widget: intensityRange},
			{Text: l.T(l.MsgDaysBetweenMed), Widget: daysBetweenMedRange},
			{Text: l.T(l.MsgDurationHours), Widget: durationHoursRange}, // New form item
			{Text: l.T(l.MsgMonths), Widget: monthsGrid},
			{Text: l.T(l.MsgName), Widget: nameEntry},
			{Text: l.T(l.MsgMedicationA), Widget: medicationAEntry},
			{Text: l.T(l.MsgMedicationB), Widget: medicationBEntry},
			{Text: l.T(l.MsgMedicationC), Widget: medicationCEntry},
			{Text: l.T(l.MsgOutputFilePath), Widget: outputFilePathContainer},
		},
	}

	// Create start button
	startButton := widget.NewButton(l.T(l.MsgStart), func() {
		// Validate intensity values
		minIntensity, minErr := validateIntensity(minIntensityEntry.Text)
		maxIntensity, maxErr := validateIntensity(maxIntensityEntry.Text)

		if minErr != nil || maxErr != nil {
			dialog.ShowError(fmt.Errorf(l.T(l.MsgIntensityError)), w)
			return
		}

		if minIntensity > maxIntensity {
			dialog.ShowError(fmt.Errorf(l.T(l.MsgMinIntensityError)), w)
			return
		}

		// Validate medication delta time
		minDaysBetweenMedication, minDaysErr := parseInt(minDaysBetweenMedicationEntry.Text)
		maxDaysBetweenMedication, maxDaysErr := parseInt(maxDaysBetweenMedicationEntry.Text)

		if minDaysErr != nil || maxDaysErr != nil {
			dialog.ShowError(fmt.Errorf(l.T(l.MsgDaysBetweenMedError)), w)
			return
		}

		if minDaysBetweenMedication > maxDaysBetweenMedication {
			dialog.ShowError(fmt.Errorf(l.T(l.MsgMinDaysBetweenMedError)), w)
			return
		}

		// Validate duration hours
		minDurationHours, minDurationErr := parseInt(minDurationHoursEntry.Text)
		maxDurationHours, maxDurationErr := parseInt(maxDurationHoursEntry.Text)

		if minDurationErr != nil || maxDurationErr != nil {
			dialog.ShowError(fmt.Errorf(l.T(l.MsgDurationHoursError)), w)
			return
		}

		if minDurationHours < 0 {
			dialog.ShowError(fmt.Errorf(l.T(l.MsgMinDurationHoursNegativeError)), w)
			return
		}

		if maxDurationHours > 24 {
			dialog.ShowError(fmt.Errorf(l.T(l.MsgMaxDurationHoursExceededError)), w)
			return
		}

		if minDurationHours > maxDurationHours {
			dialog.ShowError(fmt.Errorf(l.T(l.MsgMinDurationHoursGreaterError)), w)
			return
		}

		log.Printf("Selected months before creating config: %v\n", selectedMonths)

		config := Config{
			SampleData:               sampleDataCheck.Checked,
			MinDaysBetweenMedication: minDaysBetweenMedication,
			MaxDaysBetweenMedication: maxDaysBetweenMedication,
			MinDurationHours:         minDurationHours, // New field
			MaxDurationHours:         maxDurationHours, // New field
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

		log.Printf("Config months after creation: %v\n", config.Months)

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
	exitButton := widget.NewButton(l.T(l.MsgExit), func() {
		a.Quit()
	})

	// Create about button
	aboutButton := widget.NewButton(l.T(l.MsgAbout), func() {
		showAboutDialog(w)
	})

	// Declare content variable before using it
	var content *fyne.Container

	// Create a function to update all UI elements
	updateUI := func() {
		w.SetTitle(l.T(l.MsgAppTitle))
		sampleDataCheck.Text = l.T(l.MsgSampleData)
		sampleDataCheck.Refresh() // Refresh the sampleDataCheck widget
		// Update all other widget texts
		form.Items[0].Text = l.T(l.MsgSampleData)
		form.Items[1].Text = l.T(l.MsgIntensity)
		form.Items[2].Text = l.T(l.MsgDaysBetweenMed)
		form.Items[3].Text = l.T(l.MsgDurationHours) // Update duration hours label
		form.Items[4].Text = l.T(l.MsgMonths)
		form.Items[5].Text = l.T(l.MsgName)
		form.Items[6].Text = l.T(l.MsgMedicationA)
		form.Items[7].Text = l.T(l.MsgMedicationB)
		form.Items[8].Text = l.T(l.MsgMedicationC)
		form.Items[9].Text = l.T(l.MsgOutputFilePath)

		nameEntry.SetPlaceHolder(l.T(l.MsgNamePlaceholder))
		medicationAEntry.SetPlaceHolder(l.T(l.MsgMedicationAPlaceholder))
		medicationBEntry.SetPlaceHolder(l.T(l.MsgMedicationBPlaceholder))
		medicationCEntry.SetPlaceHolder(l.T(l.MsgMedicationCPlaceholder))
		outputFilePathEntry.SetPlaceHolder(l.T(l.MsgOutputFilePathPlaceholder))
		browseButton.SetText(l.T(l.MsgBrowse))
		startButton.SetText(l.T(l.MsgStart))
		exitButton.SetText(l.T(l.MsgExit))
		aboutButton.SetText(l.T(l.MsgAbout))

		// Update months in the checkgroup
		updateMonthsSelection(monthsGrid)
		if content != nil {
			content.Refresh()
		}
	}

	// Create language selector
	langSelector := createLanguageSelector(updateUI)

	// Create button container
	buttons := container.NewHBox(startButton, exitButton, aboutButton, langSelector)

	// Create main container
	content = container.NewVBox(
		form,
		buttons,
	)

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
	openButton := widget.NewButton(l.T(l.MsgOpenFile), func() {
		openFile(filePath)
	})

	content := container.NewVBox(
		widget.NewLabel(l.T(l.MsgSuccessGenerated)),
		widget.NewLabel(l.T(l.MsgFileSavedAt, filePath)),
		container.NewCenter(openButton),
	)

	dialog.ShowCustom(l.T(l.MsgSuccess), l.T(l.MsgClose), content, w)
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
		fmt.Printf(l.T(l.MsgErrorOpeningFile), err)
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
		return 0, fmt.Errorf(l.T(l.MsgIntensityError))
	}
	return i, nil
}

// Helper function to create the months selection widget
func createMonthsSelection(selectedMonths *[]string) *fyne.Container {
	months := []string{
		l.T(l.MsgJanuary), l.T(l.MsgFebruary), l.T(l.MsgMarch), l.T(l.MsgApril), l.T(l.MsgMay), l.T(l.MsgJune),
		l.T(l.MsgJuly), l.T(l.MsgAugust), l.T(l.MsgSeptember), l.T(l.MsgOctober), l.T(l.MsgNovember), l.T(l.MsgDecember),
	}

	checkGroup := widget.NewCheckGroup(months, func(selected []string) {
		*selectedMonths = selected
		log.Printf("Updated selected months: %v\n", *selectedMonths)
	})

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

// Helper function to update the months selection widget
func updateMonthsSelection(grid *fyne.Container) {
	months := []string{
		l.T(l.MsgJanuary), l.T(l.MsgFebruary), l.T(l.MsgMarch), l.T(l.MsgApril), l.T(l.MsgMay), l.T(l.MsgJune),
		l.T(l.MsgJuly), l.T(l.MsgAugust), l.T(l.MsgSeptember), l.T(l.MsgOctober), l.T(l.MsgNovember), l.T(l.MsgDecember),
	}

	for i, child := range grid.Objects {
		if check, ok := child.(*widget.Check); ok {
			check.Text = months[i]
			check.Refresh()
		}
	}
}

func showAboutDialog(w fyne.Window) {
	currentVersion := GetCurrentVersion()

	checkUpdatesButton := widget.NewButton(l.T(l.MsgCheckUpdates), func() {
		checkForUpdates(w)
	})

	repoURL := GithubRepoURL
	repoLink := widget.NewRichTextFromMarkdown(fmt.Sprintf("%s: [%s](%s)",
		l.T(l.MsgGitHubRepository), repoURL, repoURL))

	// Combine multiple lines into a single label
	infoText := fmt.Sprintf("%s\n%s\n%s: %s\n%s",
		l.T(l.MsgAboutTitle),
		l.T(l.MsgAboutDescription),
		l.T(l.MsgAuthor), "HugeFrog24",
		l.T(l.MsgVersion, currentVersion.String()),
	)
	infoLabel := widget.NewLabel(infoText)
	infoLabel.Wrapping = fyne.TextWrapWord

	content := container.NewVBox(
		infoLabel,
		repoLink,
		checkUpdatesButton,
	)

	dialog.ShowCustom(l.T(l.MsgAbout), l.T(l.MsgClose), content, w)
}

// Helper function to parse URL
func parseURL(urlStr string) *url.URL {
	link, err := url.Parse(urlStr)
	if err != nil {
		log.Printf("Error parsing URL: %v", err)
	}
	return link
}

func checkForUpdates(w fyne.Window) {
	log.Println("Checking for updates...")

	// Create and show a progress dialog with a cancel button
	progressBar := widget.NewProgressBarInfinite()
	cancelButton := widget.NewButton(l.T(l.MsgCancel), nil)
	content := container.NewVBox(
		widget.NewLabel(l.T(l.MsgPleaseWait)),
		progressBar,
		cancelButton,
	)
	// Create the progress dialog without any default buttons
	progressDialog := dialog.NewCustomWithoutButtons(l.T(l.MsgCheckingForUpdates), content, w)
	progressDialog.SetOnClosed(func() {
		// Add cancellation logic here if needed
	})
	progressDialog.Show()

	// Create a channel to signal cancellation
	cancelChan := make(chan struct{})

	cancelButton.OnTapped = func() {
		close(cancelChan)
		progressDialog.Hide()
	}

	// Use a goroutine to perform the update check
	go func() {
		latestVersion, downloadURL, err := CheckForUpdates(cancelChan)

		// Hide the progress dialog when done
		progressDialog.Hide()

		if err != nil {
			if err == ErrUpdateCancelled {
				log.Println("Update check cancelled by user")
				return
			}
			log.Printf("Error checking for updates: %v", err)
			dialog.ShowError(err, w)
			return
		}

		currentVersion := GetCurrentVersion()
		log.Printf("Current version: %s, latest version: %s", currentVersion, latestVersion)

		if latestVersion.GT(currentVersion) {
			log.Println("Update available")
			content := widget.NewLabel(l.T(l.MsgUpdateAvailable, latestVersion.String()))
			downloadButton := widget.NewButton(l.T(l.MsgDownload), func() {
				log.Println("Starting update download")
				progressBar := widget.NewProgressBar()
				cancelButton := widget.NewButton(l.T(l.MsgCancel), nil)
				content := container.NewVBox(
					widget.NewLabel(l.T(l.MsgPleaseWait)), // Change to "Please wait"
					progressBar,
					cancelButton,
				)
				progressDialog := dialog.NewCustomWithoutButtons(l.T(l.MsgDownloadingUpdate), content, w)
				progressDialog.Show()

				// Create a new channel for download cancellation
				downloadCancelChan := make(chan struct{})

				cancelButton.OnTapped = func() {
					close(downloadCancelChan)
					progressDialog.Hide()
				}

				go func() {
					err := DownloadAndInstallUpdate(downloadURL, func(progress float64) {
						log.Printf("Download progress: %.2f%%", progress*100)
						progressBar.SetValue(progress)
					}, downloadCancelChan)
					progressDialog.Hide()
					if err != nil {
						if err == ErrUpdateCancelled {
							log.Println("Update download cancelled by user")
							return
						}
						log.Printf("Error during update: %v", err)
						dialog.ShowError(err, w)
					} else {
						log.Println("Update completed successfully")
						restartDialog := dialog.NewConfirm(
							l.T(l.MsgUpdateSuccess),
							l.T(l.MsgRestartRequired),
							func(restart bool) {
								if restart {
									log.Println("Restarting application...")
									executable, _ := os.Executable()
									cmd := exec.Command(executable)
									cmd.Start()
									os.Exit(0)
								} else {
									log.Println("Restart postponed")
								}
							},
							w,
						)
						restartDialog.SetDismissText(l.T(l.MsgLater))
						restartDialog.SetConfirmText(l.T(l.MsgRestartNow))
						restartDialog.Show()
					}
				}()
			})
			// Add a cancel button
			cancelButton := widget.NewButton(l.T(l.MsgCancel), nil)

			buttonsContainer := container.NewHBox(downloadButton, cancelButton)

			updateDialog := dialog.NewCustomWithoutButtons(
				l.T(l.MsgUpdateAvailable, latestVersion.String()),
				container.NewVBox(content, buttonsContainer),
				w,
			)

			// Set the cancel button's OnTapped function to close the dialog
			cancelButton.OnTapped = func() {
				updateDialog.Hide()
			}

			updateDialog.Show()
		} else {
			log.Println("No updates available")
			dialog.ShowInformation(l.T(l.MsgNoUpdates), l.T(l.MsgLatestVersion), w)
		}
	}()
}

func init() {
	// Set the initial language (e.g., German)
	l.SetLanguage(language.German)
	// Set log flags to include filename and line number
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}
