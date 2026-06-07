package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Portugal — Portuguese Republic.
var dataPortugal = &Country{
	alpha2:       "PT",
	alpha3:       "PRT",
	numeric:      620,
	callingCodes: []string{"+351"},
	timezones:    []string{
		"Europe/Lisbon",
		"Atlantic/Madeira",
		"Atlantic/Azores",
	},
	tlds:         []string{".pt"},
	languages:    []xlanguage.Tag{xlanguage.Portuguese},
	currency:     currency.Eur,
	continent:    "EU",
	region:       "Europe",
	subregion:    "Southern Europe",
	flagEmoji:    "🇵🇹",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataPortugal) }
