package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Zambia — Republic of Zambia.
var dataZambia = &Country{
	alpha2:       "ZM",
	alpha3:       "ZMB",
	numeric:      894,
	callingCodes: []string{"+260"},
	timezones:    []string{"Africa/Lusaka"},
	tlds:         []string{".zm"},
	languages:    []xlanguage.Tag{xlanguage.English},
	currency:     currency.Zmw,
	continent:    "AF",
	region:       "Africa",
	subregion:    "Eastern Africa",
	flagEmoji:    "🇿🇲",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataZambia) }
