package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Togo — Togolese Republic.
var dataTogo = &Country{
	alpha2:       "TG",
	alpha3:       "TGO",
	numeric:      768,
	callingCodes: []string{"+228"},
	timezones:    []string{"Africa/Lome"},
	tlds:         []string{".tg"},
	languages:    []xlanguage.Tag{xlanguage.French},
	currency:     currency.Xof,
	continent:    "AF",
	region:       "Africa",
	subregion:    "Western Africa",
	flagEmoji:    "🇹🇬",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataTogo) }
