//go:build country_all || country_americas || country_aw || country_caribbean

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Aruba — Aruba — constituent country of the Kingdom of the Netherlands.
var dataAruba = &Country{
	alpha2:       "AW",
	alpha3:       "ABW",
	numeric:      533,
	callingCodes: []string{"+297"},
	timezones:    []string{"America/Aruba"},
	tlds:         []string{".aw"},
	officialLanguage:  xlanguage.Dutch,
	spokenLanguages:   []xlanguage.Tag{xlanguage.Dutch},
	currency:     currency.AWG,
	region:       RegionCaribbean,
	flagEmoji:    "🇦🇼",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataAruba) }

var Aruba = dataAruba
