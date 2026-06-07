//go:build country_africa || country_all || country_eastern_africa || country_yt

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Mayotte — Mayotte — overseas region of France.
var dataMayotte = &Country{
	alpha2:       "YT",
	alpha3:       "MYT",
	numeric:      175,
	callingCodes: []string{"+262"},
	timezones:    []string{"Indian/Mayotte"},
	tlds:         []string{".yt"},
	officialLanguage:  xlanguage.French,
	spokenLanguages:   []xlanguage.Tag{xlanguage.French},
	currency:     currency.EUR,
	region:       RegionEasternAfrica,
	flagEmoji:    "🇾🇹",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataMayotte) }

var Mayotte = dataMayotte
