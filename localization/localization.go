package localization

import (
	"log"
	"sync"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

var (
	printer     *message.Printer
	currentLang language.Tag
	mu          sync.Mutex
)

func init() {
	// Set the log flags before anything logs: this package's own SetLanguage call
	// below is the first log line the program emits, and without this it would be
	// the only one missing the file:line prefix.
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	// Initialize with default language (e.g., English)
	SetLanguage(language.English)
}

func SetLanguage(lang language.Tag) {
	mu.Lock()
	defer mu.Unlock()

	if printer != nil && currentLang == lang {
		return
	}

	printer = message.NewPrinter(lang)
	currentLang = lang
	log.Printf("Language set to: %s", lang)
}

func T(key string, args ...interface{}) string {
	return printer.Sprintf(key, args...)
}
