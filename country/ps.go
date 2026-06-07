//go:build country_all || country_asia || country_ps || country_western_asia

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Palestine — State of Palestine.
var dataPalestine = &Country{
	alpha2:       "PS",
	alpha3:       "PSE",
	numeric:      275,
	callingCodes: []string{"+970"},
	timezones:    []string{
		"Asia/Gaza",
		"Asia/Hebron",
	},
	tlds:         []string{
		".ps",
		".فلسطين",
	},
	officialLanguage:  xlanguage.Arabic,
	spokenLanguages:   []xlanguage.Tag{xlanguage.Arabic},
	currency:     currency.ILS,
	region:       RegionWesternAsia,
	flagEmoji:    "🇵🇸",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataPalestine) }

var Palestine = dataPalestine
