//go:build country_all || country_asia || country_iq || country_western_asia

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Iraq — Republic of Iraq.
var dataIraq = &Country{
	alpha2:       "IQ",
	alpha3:       "IRQ",
	numeric:      368,
	callingCodes: []string{"+964"},
	timezones:    []string{"Asia/Baghdad"},
	tlds:         []string{".iq"},
	officialLanguage:  xlanguage.Arabic,
	spokenLanguages:   []xlanguage.Tag{xlanguage.Arabic, xlanguage.MustParse("ku")},
	currency:     currency.IQD,
	region:       RegionWesternAsia,
	flagEmoji:    "🇮🇶",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataIraq) }

var Iraq = dataIraq
