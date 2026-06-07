//go:build country_all || country_micronesia || country_oceania || country_pw

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Palau — Republic of Palau.
var dataPalau = &Country{
	alpha2:       "PW",
	alpha3:       "PLW",
	numeric:      585,
	callingCodes: []string{"+680"},
	timezones:    []string{"Pacific/Palau"},
	tlds:         []string{".pw"},
	officialLanguage:  xlanguage.English,
	spokenLanguages:   []xlanguage.Tag{xlanguage.English, xlanguage.MustParse("pau")},
	currency:     currency.USD,
	region:       RegionMicronesia,
	flagEmoji:    "🇵🇼",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataPalau) }

var Palau = dataPalau
