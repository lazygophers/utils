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
	languages:    []xlanguage.Tag{xlanguage.Kazakh, xlanguage.Russian},
	currency:     currency.Kzt,
	continent:    "AS",
	region:       "Asia",
	subregion:    "Central Asia",
	flagEmoji:    "🇰🇿",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataKazakhstan) }
