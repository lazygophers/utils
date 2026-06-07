//go:build country_all || country_americas || country_bq || country_caribbean

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// BonaireSintEustatiusAndSaba — Caribbean Netherlands — Bonaire, Sint Eustatius and Saba.
var dataBonaireSintEustatiusAndSaba = &Country{
	alpha2:       "BQ",
	alpha3:       "BES",
	numeric:      535,
	callingCodes: []string{"+599"},
	timezones:    []string{"America/Kralendijk"},
	tlds:         []string{
		".bq",
		".nl",
	},
	languages:    []xlanguage.Tag{xlanguage.Dutch},
	currency:     currency.Usd,
	continent:    "NA",
	region:       "Americas",
	subregion:    "Caribbean",
	flagEmoji:    "🇧🇶",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataBonaireSintEustatiusAndSaba) }

var BonaireSintEustatiusAndSaba = dataBonaireSintEustatiusAndSaba
