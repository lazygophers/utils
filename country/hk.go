package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Hong Kong Special Administrative Region — independent ISO 3166-1 entry
// with its own calling code, currency, and TLD.
var dataHongKong = &Country{
	alpha2:       "HK",
	alpha3:       "HKG",
	numeric:      344,
	callingCodes: []string{"+852"},
	timezones:    []string{"Asia/Hong_Kong"},
	tlds:         []string{".hk", ".香港"},
	languages:    []xlanguage.Tag{xlanguage.MustParse("zh-Hant"), xlanguage.English},
	currency:     currency.Hkd,
	continent:    "AS",
	region:       "Asia",
	subregion:    "Eastern Asia",
	flagEmoji:    "🇭🇰",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataHongKong) }
