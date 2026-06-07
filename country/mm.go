//go:build country_all || country_asia || country_mm || country_south_eastern_asia

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Myanmar — Republic of the Union of Myanmar.
var dataMyanmar = &Country{
	alpha2:       "MM",
	alpha3:       "MMR",
	numeric:      104,
	callingCodes: []string{"+95"},
	timezones:    []string{"Asia/Yangon"},
	tlds:         []string{".mm"},
	officialLanguage:  xlanguage.MustParse("my"),
	spokenLanguages:   []xlanguage.Tag{xlanguage.MustParse("my")},
	currency:     currency.MMK,
	region:       RegionSouthEasternAsia,
	flagEmoji:    "🇲🇲",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataMyanmar) }

var Myanmar = dataMyanmar
