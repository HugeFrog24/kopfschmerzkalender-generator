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
	// Initialize with default language (e.g., English)
	SetLanguage(language.English)
	// Set log flags to include filename and line number
	log.SetFlags(log.LstdFlags | log.Lshortfile)
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
