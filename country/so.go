//go:build country_africa || country_all || country_eastern_africa || country_so

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Somalia — Federal Republic of Somalia.
var dataSomalia = &Country{
	alpha2:       "SO",
	alpha3:       "SOM",
	numeric:      706,
	callingCodes: []string{"+252"},
	timezones:    []string{"Africa/Mogadishu"},
	tlds:         []string{".so"},
	languages:    []xlanguage.Tag{xlanguage.MustParse("so"), xlanguage.Arabic},
	currency:     currency.Sos,
	continent:    "AF",
	region:       "Africa",
	subregion:    "Eastern Africa",
	flagEmoji:    "🇸🇴",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataSomalia) }

var Somalia = dataSomalia
