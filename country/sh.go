package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// SaintHelena — Saint Helena, Ascension and Tristan da Cunha — British Overseas Territory.
var dataSaintHelena = &Country{
	alpha2:       "SH",
	alpha3:       "SHN",
	numeric:      654,
	callingCodes: []string{"+290"},
	timezones:    []string{"Atlantic/St_Helena"},
	tlds:         []string{".sh"},
	languages:    []xlanguage.Tag{xlanguage.English},
	currency:     currency.Shp,
	continent:    "AF",
	region:       "Africa",
	subregion:    "Western Africa",
	flagEmoji:    "🇸🇭",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataSaintHelena) }
