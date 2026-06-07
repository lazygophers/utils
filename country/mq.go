//go:build country_all || country_americas || country_caribbean || country_mq

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Martinique — Martinique — overseas region of France.
var dataMartinique = &Country{
	alpha2:       "MQ",
	alpha3:       "MTQ",
	numeric:      474,
	callingCodes: []string{"+596"},
	timezones:    []string{"America/Martinique"},
	tlds:         []string{".mq"},
	officialLanguage:  xlanguage.French,
	spokenLanguages:   []xlanguage.Tag{xlanguage.French},
	currency:     currency.EUR,
	region:       RegionCaribbean,
	flagEmoji:    "🇲🇶",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataMartinique) }

var Martinique = dataMartinique
