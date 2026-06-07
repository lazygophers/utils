//go:build country_all || country_asia || country_pk || country_southern_asia

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Pakistan — Islamic Republic of Pakistan.
var dataPakistan = &Country{
	alpha2:       "PK",
	alpha3:       "PAK",
	numeric:      586,
	callingCodes: []string{"+92"},
	timezones:    []string{"Asia/Karachi"},
	tlds:         []string{
		".pk",
		".پاکستان",
	},
	officialLanguage:  xlanguage.Urdu,
	spokenLanguages:   []xlanguage.Tag{xlanguage.Urdu, xlanguage.English},
	currency:     currency.PKR,
	region:       RegionSouthernAsia,
	flagEmoji:    "🇵🇰",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataPakistan) }

var Pakistan = dataPakistan
