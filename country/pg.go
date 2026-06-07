//go:build country_all || country_melanesia || country_oceania || country_pg

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// PapuaNewGuinea — Independent State of Papua New Guinea.
var dataPapuaNewGuinea = &Country{
	alpha2:       "PG",
	alpha3:       "PNG",
	numeric:      598,
	callingCodes: []string{"+675"},
	timezones:    []string{
		"Pacific/Port_Moresby",
		"Pacific/Bougainville",
	},
	tlds:         []string{".pg"},
	officialLanguage:  xlanguage.English,
	spokenLanguages:   []xlanguage.Tag{xlanguage.English, xlanguage.MustParse("tpi"), xlanguage.MustParse("ho")},
	currency:     currency.PGK,
	region:       RegionMelanesia,
	flagEmoji:    "🇵🇬",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataPapuaNewGuinea) }

var PapuaNewGuinea = dataPapuaNewGuinea
