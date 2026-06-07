package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Bahrain — Kingdom of Bahrain.
var dataBahrain = &Country{
	alpha2:       "BH",
	alpha3:       "BHR",
	numeric:      48,
	callingCodes: []string{"+973"},
	timezones:    []string{"Asia/Bahrain"},
	tlds:         []string{".bh"},
	languages:    []xlanguage.Tag{xlanguage.Arabic},
	currency:     currency.Bhd,
	continent:    "AS",
	region:       "Asia",
	subregion:    "Western Asia",
	flagEmoji:    "🇧🇭",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataBahrain) }
