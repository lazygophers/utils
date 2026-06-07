//go:build country_all || country_asia || country_il || country_western_asia

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Israel — State of Israel.
var dataIsrael = &Country{
	alpha2:       "IL",
	alpha3:       "ISR",
	numeric:      376,
	callingCodes: []string{"+972"},
	timezones:    []string{"Asia/Jerusalem"},
	tlds:         []string{".il"},
	officialLanguage:  xlanguage.Hebrew,
	spokenLanguages:   []xlanguage.Tag{xlanguage.Hebrew, xlanguage.Arabic},
	currency:     currency.ILS,
	region:       RegionWesternAsia,
	flagEmoji:    "🇮🇱",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataIsrael) }

var Israel = dataIsrael
