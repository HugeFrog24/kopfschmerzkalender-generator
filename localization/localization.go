package localization

import (
	"log"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

var (
	printer *message.Printer
)

func init() {
	// Initialize with default language (e.g., English)
	SetLanguage(language.English)
	// Set log flags to include filename and line number
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func SetLanguage(lang language.Tag) {
	printer = message.NewPrinter(lang)
	log.Printf("Language set to: %s", lang)
}

func T(key string, args ...interface{}) string {
	return printer.Sprintf(key, args...)
}
