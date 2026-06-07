package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Mauritania — Islamic Republic of Mauritania.
var dataMauritania = &Country{
	alpha2:       "MR",
	alpha3:       "MRT",
	numeric:      478,
	callingCodes: []string{"+222"},
	timezones:    []string{"Africa/Nouakchott"},
	tlds:         []string{".mr"},
	languages:    []xlanguage.Tag{xlanguage.Arabic},
	currency:     currency.Mru,
	continent:    "AF",
	region:       "Africa",
	subregion:    "Western Africa",
	flagEmoji:    "🇲🇷",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataMauritania) }
