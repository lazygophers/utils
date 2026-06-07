//go:build country_africa || country_all || country_ml || country_western_africa

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Mali — Republic of Mali.
var dataMali = &Country{
	alpha2:       "ML",
	alpha3:       "MLI",
	numeric:      466,
	callingCodes: []string{"+223"},
	timezones:    []string{"Africa/Bamako"},
	tlds:         []string{".ml"},
	officialLanguage:  xlanguage.French,
	spokenLanguages:   []xlanguage.Tag{xlanguage.French},
	currency:     currency.XOF,
	region:       RegionWesternAfrica,
	flagEmoji:    "🇲🇱",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataMali) }

var Mali = dataMali
