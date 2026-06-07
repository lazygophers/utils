package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// CaboVerde — Republic of Cabo Verde.
var dataCaboVerde = &Country{
	alpha2:       "CV",
	alpha3:       "CPV",
	numeric:      132,
	callingCodes: []string{"+238"},
	timezones:    []string{"Atlantic/Cape_Verde"},
	tlds:         []string{".cv"},
	languages:    []xlanguage.Tag{xlanguage.Portuguese},
	currency:     currency.Cve,
	continent:    "AF",
	region:       "Africa",
	subregion:    "Western Africa",
	flagEmoji:    "🇨🇻",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataCaboVerde) }
