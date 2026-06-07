//go:build country_all || country_asia || country_central_asia || country_kz

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Kazakhstan — Republic of Kazakhstan.
var dataKazakhstan = &Country{
	alpha2:       "KZ",
	alpha3:       "KAZ",
	numeric:      398,
	callingCodes: []string{"+7"},
	timezones:    []string{
		"Asia/Almaty",
		"Asia/Aqtau",
		"Asia/Aqtobe",
		"Asia/Atyrau",
		"Asia/Oral",
		"Asia/Qostanay",
		"Asia/Qyzylorda",
	},
	tlds:         []string{".kz"},
	officialLanguage:  xlanguage.Kazakh,
	spokenLanguages:   []xlanguage.Tag{xlanguage.Kazakh, xlanguage.Russian},
	currency:     currency.KZT,
	region:       RegionCentralAsia,
	flagEmoji:    "🇰🇿",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataKazakhstan) }

var Kazakhstan = dataKazakhstan
