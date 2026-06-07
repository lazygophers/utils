//go:build country_all || country_asia || country_bt || country_southern_asia

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Bhutan — Kingdom of Bhutan.
var dataBhutan = &Country{
	alpha2:       "BT",
	alpha3:       "BTN",
	numeric:      64,
	callingCodes: []string{"+975"},
	timezones:    []string{"Asia/Thimphu"},
	tlds:         []string{".bt"},
	languages:    []xlanguage.Tag{xlanguage.MustParse("dz")},
	currency:     currency.Btn,
	continent:    "AS",
	region:       "Asia",
	subregion:    "Southern Asia",
	flagEmoji:    "🇧🇹",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataBhutan) }

var Bhutan = dataBhutan
