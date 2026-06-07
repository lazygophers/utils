package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Singapore — Republic of Singapore.
var dataSingapore = &Country{
	alpha2:       "SG",
	alpha3:       "SGP",
	numeric:      702,
	callingCodes: []string{"+65"},
	timezones:    []string{"Asia/Singapore"},
	tlds:         []string{
		".sg",
		".新加坡",
		".சிங்கப்பூர்",
	},
	languages:    []xlanguage.Tag{xlanguage.English, xlanguage.Chinese, xlanguage.Malay, xlanguage.MustParse("ta")},
	currency:     currency.Sgd,
	continent:    "AS",
	region:       "Asia",
	subregion:    "South-eastern Asia",
	flagEmoji:    "🇸🇬",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataSingapore) }
