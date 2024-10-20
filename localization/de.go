package localization

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func init() {
	// German translations
	message.SetString(language.German, MsgAppTitle, "Kopfschmerzkalender-Generator")
	message.SetString(language.German, MsgSampleData, "Beispieldaten")
	message.SetString(language.German, MsgMinIntensity, "Min. Intensität (0-10)")
	message.SetString(language.German, MsgMaxIntensity, "Max. Intensität (0-10)")
	message.SetString(language.German, MsgMinDaysBetweenMed, "Min. Tage zwischen Medikation")
	message.SetString(language.German, MsgMaxDaysBetweenMed, "Max. Tage zwischen Medikation")
	message.SetString(language.German, MsgMonths, "Monate")
	message.SetString(language.German, MsgName, "Name")
	message.SetString(language.German, MsgMedicationA, "Medikament A")
	message.SetString(language.German, MsgMedicationB, "Medikament B")
	message.SetString(language.German, MsgMedicationC, "Medikament C")
	message.SetString(language.German, MsgOutputFilePath, "Ausgabedateipfad")
	message.SetString(language.German, MsgBrowse, "Durchsuchen")
	message.SetString(language.German, MsgStart, "Start")
	message.SetString(language.German, MsgExit, "Beenden")
	message.SetString(language.German, MsgNamePlaceholder, "Max Mustermann")
	message.SetString(language.German, MsgMedicationAPlaceholder, "z.B., Ibuprofen 800")
	message.SetString(language.German, MsgMedicationBPlaceholder, "z.B., Thomapyrin Duo")
	message.SetString(language.German, MsgMedicationCPlaceholder, "z.B. Naproxen 500")
	message.SetString(language.German, MsgOutputFilePathPlaceholder, "Leer lassen für Standardpfad")
	message.SetString(language.German, MsgOpenFile, "Datei öffnen")
	message.SetString(language.German, MsgSuccessGenerated, "Kopfschmerzkalender erfolgreich generiert")
	message.SetString(language.German, MsgFileSavedAt, "Datei gespeichert unter: %s")
	message.SetString(language.German, MsgSuccess, "Erfolg")
	message.SetString(language.German, MsgClose, "Schließen")
	message.SetString(language.German, MsgJanuary, "Januar")
	message.SetString(language.German, MsgFebruary, "Februar")
	message.SetString(language.German, MsgMarch, "März")
	message.SetString(language.German, MsgApril, "April")
	message.SetString(language.German, MsgMay, "Mai")
	message.SetString(language.German, MsgJune, "Juni")
	message.SetString(language.German, MsgJuly, "Juli")
	message.SetString(language.German, MsgAugust, "August")
	message.SetString(language.German, MsgSeptember, "September")
	message.SetString(language.German, MsgOctober, "Oktober")
	message.SetString(language.German, MsgNovember, "November")
	message.SetString(language.German, MsgDecember, "Dezember")
	message.SetString(language.German, MsgIntensityError, "Ungültiger Intensitätswert. Bitte geben Sie eine Zahl zwischen 0 und 10 ein.")
	message.SetString(language.German, MsgMinIntensityError, "Die minimale Intensität darf nicht größer sein als die maximale Intensität.")
	message.SetString(language.German, MsgDaysBetweenMedError, "Ungültige Anzahl von Tagen zwischen Medikationen. Bitte geben Sie eine gültige Zahl ein.")
	message.SetString(language.German, MsgMinDaysBetweenMedError, "Die minimale Anzahl von Tagen zwischen Medikationen darf nicht größer sein als die maximale Anzahl.")
	message.SetString(language.German, MsgAbout, "Über")
	message.SetString(language.German, MsgAboutTitle, "Kopfschmerzkalender-Generator")
	message.SetString(language.German, MsgAboutDescription, "Ein Programm zur Erstellung von Kopfschmerzkalendern.")
	message.SetString(language.German, MsgAuthor, "Autor")
	message.SetString(language.German, MsgVersion, "Version: %s")
	message.SetString(language.German, MsgCheckUpdates, "Nach Updates suchen")
	message.SetString(language.German, MsgCheckingForUpdates, "Es wird nach Updates gesucht...")
	message.SetString(language.German, MsgUpdateAvailable, "Neue Version %s ist verfügbar!")
	message.SetString(language.German, MsgNoUpdates, "Keine Updates verfügbar")
	message.SetString(language.German, MsgLatestVersion, "Sie verwenden die neueste Version.")
	message.SetString(language.German, MsgDownload, "Herunterladen")
	message.SetString(language.German, MsgDownloadingUpdate, "Update wird heruntergeladen...")
	message.SetString(language.German, MsgPleaseWait, "Bitte warten...")
	message.SetString(language.German, MsgUpdateSuccess, "Update erfolgreich")
	message.SetString(language.German, MsgRestartRequired, "Bitte starten Sie die Anwendung neu, um das Update anzuwenden.")
	message.SetString(language.German, MsgLater, "Später")
	message.SetString(language.German, MsgRestartNow, "Jetzt neu starten")
	message.SetString(language.German, MsgCancel, "Abbrechen")
	message.SetString(language.German, MsgIntensity, "Intensität")
	message.SetString(language.German, MsgDaysBetweenMed, "Tage zwischen Medikation")
	message.SetString(language.German, MsgMinDurationHoursNegativeError, "Die minimale Dauer darf nicht negativ sein.")
	message.SetString(language.German, MsgMaxDurationHoursExceededError, "Die maximale Dauer darf nicht mehr als 24 Stunden betragen.")
	message.SetString(language.German, MsgMinDurationHoursGreaterError, "Die minimale Dauer darf nicht größer sein als die maximale Dauer.")
	message.SetString(language.German, MsgDurationHours, "Dauer (Stunden)")
	message.SetString(language.German, MsgMinDurationHours, "Min. Dauer (Stunden)")
	message.SetString(language.German, MsgMaxDurationHours, "Max. Dauer (Stunden)")
	message.SetString(language.German, MsgDurationHoursError, "Ungültige Dauer. Bitte geben Sie eine gültige Anzahl von Stunden ein.")
	message.SetString(language.German, MsgMinDurationHoursPlaceholder, "z.B. 1")
	message.SetString(language.German, MsgMaxDurationHoursPlaceholder, "z.B. 24")
	message.SetString(language.German, MsgGitHubRepository, "GitHub-Repository")
}
