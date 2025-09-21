package human

import (
	"fmt"
	"testing"
)

func TestDebugLocales(t *testing.T) {
	testCases := []string{"zh-TW", "fr", "ru", "ar", "es"}

	for _, locale := range testCases {
		config, ok := GetLocaleConfig(locale)
		if ok {
			fmt.Printf("Locale %s: Language=%s, Region=%s\n", locale, config.Language, config.Region)
		} else {
			fmt.Printf("Locale %s: NOT FOUND\n", locale)
		}
	}
}