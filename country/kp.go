package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// NorthKorea — Democratic People's Republic of Korea.
var dataNorthKorea = &Country{
	alpha2:       "KP",
	alpha3:       "PRK",
	numeric:      408,
	callingCodes: []string{"+850"},
	timezones:    []string{"Asia/Pyongyang"},
	tlds:         []string{
		".kp",
		".조선",
	},
	languages:    []xlanguage.Tag{xlanguage.Korean},
	currency:     currency.Kpw,
	continent:    "AS",
	region:       "Asia",
	subregion:    "Eastern Asia",
	flagEmoji:    "🇰🇵",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataNorthKorea) }
