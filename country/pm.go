//go:build country_all || country_americas || country_northern_america || country_pm

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// SaintPierreAndMiquelon — Saint Pierre and Miquelon — overseas collectivity of France.
var dataSaintPierreAndMiquelon = &Country{
	alpha2:       "PM",
	alpha3:       "SPM",
	numeric:      666,
	callingCodes: []string{"+508"},
	timezones:    []string{"America/Miquelon"},
	tlds:         []string{".pm"},
	officialLanguage:  xlanguage.French,
	spokenLanguages:   []xlanguage.Tag{xlanguage.French},
	currency:     currency.EUR,
	region:       RegionNorthernAmerica,
	flagEmoji:    "🇵🇲",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataSaintPierreAndMiquelon) }

var SaintPierreAndMiquelon = dataSaintPierreAndMiquelon
