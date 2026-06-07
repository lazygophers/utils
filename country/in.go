package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// India — Republic of India.
var dataIndia = &Country{
	alpha2:       "IN",
	alpha3:       "IND",
	numeric:      356,
	callingCodes: []string{"+91"},
	timezones:    []string{"Asia/Kolkata"},
	tlds:         []string{
		".in",
		".भारत",
	},
	officialLanguage:  xlanguage.Hindi,
	spokenLanguages:   []xlanguage.Tag{xlanguage.Hindi, xlanguage.English, xlanguage.Tamil, xlanguage.Bengali, xlanguage.Telugu, xlanguage.Marathi, xlanguage.Urdu, xlanguage.Gujarati, xlanguage.Kannada, xlanguage.Punjabi},
	currency:     currency.INR,
	region:       RegionSouthernAsia,
	flagEmoji:    "🇮🇳",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataIndia) }

var India = dataIndia
