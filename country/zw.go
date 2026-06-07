//go:build country_africa || country_all || country_eastern_africa || country_zw

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Zimbabwe — Republic of Zimbabwe.
var dataZimbabwe = &Country{
	alpha2:       "ZW",
	alpha3:       "ZWE",
	numeric:      716,
	callingCodes: []string{"+263"},
	timezones:    []string{"Africa/Harare"},
	tlds:         []string{".zw"},
	languages:    []xlanguage.Tag{xlanguage.English, xlanguage.MustParse("sn")},
	currency:     currency.Zwl,
	continent:    "AF",
	region:       "Africa",
	subregion:    "Eastern Africa",
	flagEmoji:    "🇿🇼",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataZimbabwe) }

var Zimbabwe = dataZimbabwe
