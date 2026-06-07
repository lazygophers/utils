//go:build country_all || country_asia || country_jo || country_western_asia

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Jordan — Hashemite Kingdom of Jordan.
var dataJordan = &Country{
	alpha2:       "JO",
	alpha3:       "JOR",
	numeric:      400,
	callingCodes: []string{"+962"},
	timezones:    []string{"Asia/Amman"},
	tlds:         []string{
		".jo",
		".الاردن",
	},
	officialLanguage:  xlanguage.Arabic,
	spokenLanguages:   []xlanguage.Tag{xlanguage.Arabic},
	currency:     currency.JOD,
	region:       RegionWesternAsia,
	flagEmoji:    "🇯🇴",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataJordan) }

var Jordan = dataJordan
