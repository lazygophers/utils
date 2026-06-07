//go:build country_all || country_americas || country_bl || country_caribbean

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// SaintBarthelemy — Collectivity of Saint Barthélemy.
var dataSaintBarthelemy = &Country{
	alpha2:       "BL",
	alpha3:       "BLM",
	numeric:      652,
	callingCodes: []string{"+590"},
	timezones:    []string{"America/St_Barthelemy"},
	tlds:         []string{".bl"},
	officialLanguage:  xlanguage.French,
	spokenLanguages:   []xlanguage.Tag{xlanguage.French},
	currency:     currency.EUR,
	region:       RegionCaribbean,
	flagEmoji:    "🇧🇱",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataSaintBarthelemy) }

var SaintBarthelemy = dataSaintBarthelemy
