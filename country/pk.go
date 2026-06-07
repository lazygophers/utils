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
	languages:    []xlanguage.Tag{xlanguage.Urdu, xlanguage.English},
	currency:     currency.Pkr,
	continent:    "AS",
	region:       "Asia",
	subregion:    "Southern Asia",
	flagEmoji:    "🇵🇰",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataPakistan) }
