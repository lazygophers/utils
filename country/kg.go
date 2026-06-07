//go:build country_all || country_asia || country_central_asia || country_kg

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Kyrgyzstan — Kyrgyz Republic.
var dataKyrgyzstan = &Country{
	alpha2:       "KG",
	alpha3:       "KGZ",
	numeric:      417,
	callingCodes: []string{"+996"},
	timezones:    []string{"Asia/Bishkek"},
	tlds:         []string{".kg"},
	officialLanguage:  xlanguage.MustParse("ky"),
	spokenLanguages:   []xlanguage.Tag{xlanguage.MustParse("ky"), xlanguage.Russian},
	currency:     currency.KGS,
	region:       RegionCentralAsia,
	flagEmoji:    "🇰🇬",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataKyrgyzstan) }

var Kyrgyzstan = dataKyrgyzstan
