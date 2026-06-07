package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// DominicanRepublic — Dominican Republic.
var dataDominicanRepublic = &Country{
	alpha2:       "DO",
	alpha3:       "DOM",
	numeric:      214,
	callingCodes: []string{
		"+1-809",
		"+1-829",
		"+1-849",
	},
	timezones:    []string{"America/Santo_Domingo"},
	tlds:         []string{".do"},
	languages:    []xlanguage.Tag{xlanguage.Spanish},
	currency:     currency.Dop,
	continent:    "NA",
	region:       "Americas",
	subregion:    "Caribbean",
	flagEmoji:    "🇩🇴",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataDominicanRepublic) }
