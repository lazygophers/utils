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
	languages:    []xlanguage.Tag{xlanguage.Arabic},
	currency:     currency.Ils,
	continent:    "AS",
	region:       "Asia",
	subregion:    "Western Asia",
	flagEmoji:    "🇵🇸",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataPalestine) }
