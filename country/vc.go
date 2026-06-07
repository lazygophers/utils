package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// SaintVincentAndGrenadines — Saint Vincent and the Grenadines.
var dataSaintVincentAndGrenadines = &Country{
	alpha2:       "VC",
	alpha3:       "VCT",
	numeric:      670,
	callingCodes: []string{"+1-784"},
	timezones:    []string{"America/St_Vincent"},
	tlds:         []string{".vc"},
	languages:    []xlanguage.Tag{xlanguage.English},
	currency:     currency.Xcd,
	continent:    "NA",
	region:       "Americas",
	subregion:    "Caribbean",
	flagEmoji:    "🇻🇨",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataSaintVincentAndGrenadines) }
