//go:build country_all || country_europe || country_mt || country_southern_europe

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Malta — Republic of Malta.
var dataMalta = &Country{
	alpha2:       "MT",
	alpha3:       "MLT",
	numeric:      470,
	callingCodes: []string{"+356"},
	timezones:    []string{"Europe/Malta"},
	tlds:         []string{".mt"},
	officialLanguage:  xlanguage.MustParse("mt"),
	spokenLanguages:   []xlanguage.Tag{xlanguage.MustParse("mt"), xlanguage.English},
	currency:     currency.EUR,
	region:       RegionSouthernEurope,
	flagEmoji:    "🇲🇹",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataMalta) }

var Malta = dataMalta
