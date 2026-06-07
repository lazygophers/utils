package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Albania — Republic of Albania.
var dataAlbania = &Country{
	alpha2:       "AL",
	alpha3:       "ALB",
	numeric:      8,
	callingCodes: []string{"+355"},
	timezones:    []string{"Europe/Tirane"},
	tlds:         []string{".al"},
	languages:    []xlanguage.Tag{xlanguage.MustParse("sq")},
	currency:     currency.All,
	continent:    "EU",
	region:       "Europe",
	subregion:    "Southern Europe",
	flagEmoji:    "🇦🇱",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataAlbania) }
