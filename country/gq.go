//go:build country_africa || country_all || country_gq || country_middle_africa

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// EquatorialGuinea — Republic of Equatorial Guinea.
var dataEquatorialGuinea = &Country{
	alpha2:       "GQ",
	alpha3:       "GNQ",
	numeric:      226,
	callingCodes: []string{"+240"},
	timezones:    []string{"Africa/Malabo"},
	tlds:         []string{".gq"},
	officialLanguage:  xlanguage.Spanish,
	spokenLanguages:   []xlanguage.Tag{xlanguage.Spanish, xlanguage.French, xlanguage.Portuguese},
	currency:     currency.XAF,
	region:       RegionMiddleAfrica,
	flagEmoji:    "🇬🇶",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataEquatorialGuinea) }

var EquatorialGuinea = dataEquatorialGuinea
