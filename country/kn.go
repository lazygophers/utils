package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// SaintKittsAndNevis — Federation of Saint Kitts and Nevis.
var dataSaintKittsAndNevis = &Country{
	alpha2:       "KN",
	alpha3:       "KNA",
	numeric:      659,
	callingCodes: []string{"+1-869"},
	timezones:    []string{"America/St_Kitts"},
	tlds:         []string{".kn"},
	languages:    []xlanguage.Tag{xlanguage.English},
	currency:     currency.Xcd,
	continent:    "NA",
	region:       "Americas",
	subregion:    "Caribbean",
	flagEmoji:    "🇰🇳",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataSaintKittsAndNevis) }
