//go:build country_all || country_asia || country_id || country_south_eastern_asia

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Indonesia — Republic of Indonesia.
var dataIndonesia = &Country{
	alpha2:       "ID",
	alpha3:       "IDN",
	numeric:      360,
	callingCodes: []string{"+62"},
	timezones:    []string{
		"Asia/Jakarta",
		"Asia/Makassar",
		"Asia/Jayapura",
		"Asia/Pontianak",
	},
	tlds:         []string{".id"},
	officialLanguage:  xlanguage.Indonesian,
	spokenLanguages:   []xlanguage.Tag{xlanguage.Indonesian},
	currency:     currency.IDR,
	region:       RegionSouthEasternAsia,
	flagEmoji:    "🇮🇩",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataIndonesia) }

var Indonesia = dataIndonesia
