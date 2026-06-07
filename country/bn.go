//go:build country_all || country_asia || country_bn || country_south_eastern_asia

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Brunei — Nation of Brunei, the Abode of Peace.
var dataBrunei = &Country{
	alpha2:       "BN",
	alpha3:       "BRN",
	numeric:      96,
	callingCodes: []string{"+673"},
	timezones:    []string{"Asia/Brunei"},
	tlds:         []string{".bn"},
	officialLanguage:  xlanguage.Malay,
	spokenLanguages:   []xlanguage.Tag{xlanguage.Malay},
	currency:     currency.BND,
	region:       RegionSouthEasternAsia,
	flagEmoji:    "🇧🇳",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataBrunei) }

var Brunei = dataBrunei
