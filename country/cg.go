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
	languages:    []xlanguage.Tag{xlanguage.French},
	currency:     currency.Xaf,
	continent:    "AF",
	region:       "Africa",
	subregion:    "Middle Africa",
	flagEmoji:    "🇨🇬",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataCongo) }
