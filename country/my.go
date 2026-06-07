package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Malaysia — Malaysia.
var dataMalaysia = &Country{
	alpha2:       "MY",
	alpha3:       "MYS",
	numeric:      458,
	callingCodes: []string{"+60"},
	timezones:    []string{
		"Asia/Kuala_Lumpur",
		"Asia/Kuching",
	},
	tlds:         []string{".my"},
	languages:    []xlanguage.Tag{xlanguage.Malay},
	currency:     currency.Myr,
	continent:    "AS",
	region:       "Asia",
	subregion:    "South-eastern Asia",
	flagEmoji:    "🇲🇾",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataMalaysia) }
