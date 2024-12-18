package localization

import (
	"log"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func init() {
	// English translations
	setEnglishString := func(key, value string) {
		err := message.SetString(language.English, key, value)
		if err != nil {
			log.Printf("Error setting English translation for %s: %v\n", key, err)
		}
	}

	setEnglishString(MsgAppTitle, "Kopfschmerzkalender-Generator")
	setEnglishString(MsgSampleData, "Sample Data")
	setEnglishString(MsgMinIntensity, "Min Intensity (0-10)")
	setEnglishString(MsgMaxIntensity, "Max Intensity (0-10)")
	setEnglishString(MsgMinDaysBetweenMed, "Min Days Between Medication")
	setEnglishString(MsgMaxDaysBetweenMed, "Max Days Between Medication")
	setEnglishString(MsgMonths, "Months")
	setEnglishString(MsgName, "Name")
	setEnglishString(MsgMedicationA, "Medication A")
	setEnglishString(MsgMedicationB, "Medication B")
	setEnglishString(MsgMedicationC, "Medication C")
	setEnglishString(MsgOutputFilePath, "Output File Path")
	setEnglishString(MsgBrowse, "Browse")
	setEnglishString(MsgStart, "Start")
	setEnglishString(MsgExit, "Exit")
	setEnglishString(MsgNamePlaceholder, "Max Mustermann")
	setEnglishString(MsgMedicationAPlaceholder, "e.g., Ibuprofen 800")
	setEnglishString(MsgMedicationBPlaceholder, "e.g., Thomapyrin Duo")
	setEnglishString(MsgMedicationCPlaceholder, "e.g. Naproxen 500")
	setEnglishString(MsgOutputFilePathPlaceholder, "Leave empty for default path")
	setEnglishString(MsgOpenFile, "Open File")
	setEnglishString(MsgSuccessGenerated, "Kopfschmerzkalender generated successfully")
	setEnglishString(MsgFileSavedAt, "File saved at: %s")
	setEnglishString(MsgSuccess, "Success")
	setEnglishString(MsgClose, "Close")
	setEnglishString(MsgJanuary, "January")
	setEnglishString(MsgFebruary, "February")
	setEnglishString(MsgMarch, "March")
	setEnglishString(MsgApril, "April")
	setEnglishString(MsgMay, "May")
	setEnglishString(MsgJune, "June")
	setEnglishString(MsgJuly, "July")
	setEnglishString(MsgAugust, "August")
	setEnglishString(MsgSeptember, "September")
	setEnglishString(MsgOctober, "October")
	setEnglishString(MsgNovember, "November")
	setEnglishString(MsgDecember, "December")
	setEnglishString(MsgIntensityError, "Invalid intensity value. Please enter a number between 0 and 10.")
	setEnglishString(MsgMinIntensityError, "Minimum intensity cannot be greater than maximum intensity.")
	setEnglishString(MsgDaysBetweenMedError, "Invalid number of days between medication. Please enter a valid number.")
	setEnglishString(MsgMinDaysBetweenMedError, "Minimum days between medication cannot be greater than maximum days.")
	setEnglishString(MsgAbout, "About")
	setEnglishString(MsgAboutTitle, "Kopfschmerzkalender-Generator")
	setEnglishString(MsgAboutDescription, "A tool to generate headache calendars.")
	setEnglishString(MsgAuthor, "Author")
	setEnglishString(MsgVersion, "Version: %s")
	setEnglishString(MsgCheckUpdates, "Check for Updates")
	setEnglishString(MsgCheckingForUpdates, "Checking for updates...")
	setEnglishString(MsgUpdateAvailable, "New version %s is available!")
	setEnglishString(MsgNoUpdates, "No Updates Available")
	setEnglishString(MsgLatestVersion, "You are using the latest version.")
	setEnglishString(MsgDownload, "Download")
	setEnglishString(MsgDownloadingUpdate, "Downloading Update...")
	setEnglishString(MsgPleaseWait, "Please wait...")
	setEnglishString(MsgUpdateSuccess, "Update Successful")
	setEnglishString(MsgRestartRequired, "Please restart the application to apply the update.")
	setEnglishString(MsgLater, "Later")
	setEnglishString(MsgRestartNow, "Restart Now")
	setEnglishString(MsgCancel, "Cancel")
	setEnglishString(MsgIntensity, "Intensity")
	setEnglishString(MsgDaysBetweenMed, "Days Between Medication")
	setEnglishString(MsgMinDurationHoursNegativeError, "Minimum duration cannot be negative.")
	setEnglishString(MsgMaxDurationHoursExceededError, "Maximum duration cannot exceed 24 hours.")
	setEnglishString(MsgMinDurationHoursGreaterError, "Minimum duration cannot be greater than maximum duration.")
	setEnglishString(MsgDurationHours, "Duration (Hours)")
	setEnglishString(MsgMinDurationHours, "Min Duration (Hours)")
	setEnglishString(MsgMaxDurationHours, "Max Duration (Hours)")
	setEnglishString(MsgDurationHoursError, "Invalid duration. Please enter a valid number of hours.")
	setEnglishString(MsgMinDurationHoursPlaceholder, "e.g., 1")
	setEnglishString(MsgMaxDurationHoursPlaceholder, "e.g., 24")
	setEnglishString(MsgGitHubRepository, "GitHub Repository")
}
