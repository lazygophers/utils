//go:build country_africa || country_all || country_sl || country_western_africa

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// SierraLeone — Republic of Sierra Leone.
var dataSierraLeone = &Country{
	alpha2:       "SL",
	alpha3:       "SLE",
	numeric:      694,
	callingCodes: []string{"+232"},
	timezones:    []string{"Africa/Freetown"},
	tlds:         []string{".sl"},
	officialLanguage:  xlanguage.English,
	spokenLanguages:   []xlanguage.Tag{xlanguage.English},
	currency:     currency.SLL,
	region:       RegionWesternAfrica,
	flagEmoji:    "🇸🇱",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataSierraLeone) }

var SierraLeone = dataSierraLeone
