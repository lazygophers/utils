package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Japan — Japan.
var dataJapan = &Country{
	alpha2:       "JP",
	alpha3:       "JPN",
	numeric:      392,
	callingCodes: []string{"+81"},
	timezones:    []string{"Asia/Tokyo"},
	tlds:         []string{
		".jp",
		".日本",
	},
	languages:    []xlanguage.Tag{xlanguage.Japanese},
	currency:     currency.Jpy,
	continent:    "AS",
	region:       "Asia",
	subregion:    "Eastern Asia",
	flagEmoji:    "🇯🇵",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataJapan) }
