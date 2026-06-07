//go:build country_africa || country_all || country_cg || country_middle_africa

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Congo — Republic of the Congo.
var dataCongo = &Country{
	alpha2:       "CG",
	alpha3:       "COG",
	numeric:      178,
	callingCodes: []string{"+242"},
	timezones:    []string{"Africa/Brazzaville"},
	tlds:         []string{".cg"},
	officialLanguage:  xlanguage.French,
	spokenLanguages:   []xlanguage.Tag{xlanguage.French},
	currency:     currency.XAF,
	region:       RegionMiddleAfrica,
	flagEmoji:    "🇨🇬",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataCongo) }

var Congo = dataCongo
