package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// BosniaAndHerzegovina — Bosnia and Herzegovina.
var dataBosniaAndHerzegovina = &Country{
	alpha2:       "BA",
	alpha3:       "BIH",
	numeric:      70,
	callingCodes: []string{"+387"},
	timezones:    []string{"Europe/Sarajevo"},
	tlds:         []string{".ba"},
	languages:    []xlanguage.Tag{xlanguage.MustParse("bs"), xlanguage.Croatian, xlanguage.Serbian},
	currency:     currency.Bam,
	continent:    "EU",
	region:       "Europe",
	subregion:    "Southern Europe",
	flagEmoji:    "🇧🇦",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataBosniaAndHerzegovina) }
