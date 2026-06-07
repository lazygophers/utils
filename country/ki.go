//go:build country_all || country_ki || country_micronesia || country_oceania

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Kiribati — Republic of Kiribati.
var dataKiribati = &Country{
	alpha2:       "KI",
	alpha3:       "KIR",
	numeric:      296,
	callingCodes: []string{"+686"},
	timezones:    []string{
		"Pacific/Tarawa",
		"Pacific/Kanton",
		"Pacific/Kiritimati",
	},
	tlds:         []string{".ki"},
	officialLanguage:  xlanguage.English,
	spokenLanguages:   []xlanguage.Tag{xlanguage.English, xlanguage.MustParse("gil")},
	currency:     currency.AUD,
	region:       RegionMicronesia,
	flagEmoji:    "🇰🇮",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataKiribati) }

var Kiribati = dataKiribati
