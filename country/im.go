package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// IsleOfMan — Isle of Man — British Crown Dependency.
var dataIsleOfMan = &Country{
	alpha2:       "IM",
	alpha3:       "IMN",
	numeric:      833,
	callingCodes: []string{"+44-1624"},
	timezones:    []string{"Europe/Isle_of_Man"},
	tlds:         []string{".im"},
	languages:    []xlanguage.Tag{xlanguage.English},
	currency:     currency.Gbp,
	continent:    "EU",
	region:       "Europe",
	subregion:    "Northern Europe",
	flagEmoji:    "🇮🇲",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataIsleOfMan) }
