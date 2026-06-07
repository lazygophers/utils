//go:build country_all || country_am || country_asia || country_western_asia

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Armenia — Republic of Armenia.
var dataArmenia = &Country{
	alpha2:       "AM",
	alpha3:       "ARM",
	numeric:      51,
	callingCodes: []string{"+374"},
	timezones:    []string{"Asia/Yerevan"},
	tlds:         []string{".am"},
	officialLanguage:  xlanguage.MustParse("hy"),
	spokenLanguages:   []xlanguage.Tag{xlanguage.MustParse("hy"), xlanguage.Russian},
	currency:     currency.AMD,
	region:       RegionWesternAsia,
	flagEmoji:    "🇦🇲",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataArmenia) }

var Armenia = dataArmenia
