//go:build country_all || country_au || country_australia_and_new_zealand || country_oceania

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Australia — Commonwealth of Australia.
var dataAustralia = &Country{
	alpha2:       "AU",
	alpha3:       "AUS",
	numeric:      36,
	callingCodes: []string{"+61"},
	timezones:    []string{
		"Australia/Sydney",
		"Australia/Melbourne",
		"Australia/Brisbane",
		"Australia/Perth",
		"Australia/Adelaide",
		"Australia/Hobart",
		"Australia/Darwin",
	},
	tlds:         []string{".au"},
	officialLanguage:  xlanguage.English,
	spokenLanguages:   []xlanguage.Tag{xlanguage.English},
	currency:     currency.AUD,
	region:       RegionAustraliaAndNewZealand,
	flagEmoji:    "🇦🇺",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataAustralia) }

var Australia = dataAustralia
