//go:build country_all || country_asia || country_az || country_western_asia

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Azerbaijan — Republic of Azerbaijan.
var dataAzerbaijan = &Country{
	alpha2:       "AZ",
	alpha3:       "AZE",
	numeric:      31,
	callingCodes: []string{"+994"},
	timezones:    []string{"Asia/Baku"},
	tlds:         []string{".az"},
	officialLanguage:  xlanguage.MustParse("az"),
	spokenLanguages:   []xlanguage.Tag{xlanguage.MustParse("az")},
	currency:     currency.AZN,
	region:       RegionWesternAsia,
	flagEmoji:    "🇦🇿",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataAzerbaijan) }

var Azerbaijan = dataAzerbaijan
