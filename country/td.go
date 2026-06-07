//go:build country_africa || country_all || country_middle_africa || country_td

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Chad — Republic of Chad.
var dataChad = &Country{
	alpha2:       "TD",
	alpha3:       "TCD",
	numeric:      148,
	callingCodes: []string{"+235"},
	timezones:    []string{"Africa/Ndjamena"},
	tlds:         []string{".td"},
	officialLanguage:  xlanguage.French,
	spokenLanguages:   []xlanguage.Tag{xlanguage.French, xlanguage.Arabic},
	currency:     currency.XAF,
	region:       RegionMiddleAfrica,
	flagEmoji:    "🇹🇩",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataChad) }

var Chad = dataChad
