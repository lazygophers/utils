package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// SintMaarten — Sint Maarten (Dutch part).
var dataSintMaarten = &Country{
	alpha2:       "SX",
	alpha3:       "SXM",
	numeric:      534,
	callingCodes: []string{"+1-721"},
	timezones:    []string{"America/Lower_Princes"},
	tlds:         []string{".sx"},
	languages:    []xlanguage.Tag{xlanguage.Dutch, xlanguage.English},
	currency:     currency.Ang,
	continent:    "NA",
	region:       "Americas",
	subregion:    "Caribbean",
	flagEmoji:    "🇸🇽",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataSintMaarten) }
