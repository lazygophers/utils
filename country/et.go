//go:build country_africa || country_all || country_eastern_africa || country_et

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Ethiopia — Federal Democratic Republic of Ethiopia.
var dataEthiopia = &Country{
	alpha2:       "ET",
	alpha3:       "ETH",
	numeric:      231,
	callingCodes: []string{"+251"},
	timezones:    []string{"Africa/Addis_Ababa"},
	tlds:         []string{".et"},
	officialLanguage:  xlanguage.Amharic,
	spokenLanguages:   []xlanguage.Tag{xlanguage.Amharic, xlanguage.English},
	currency:     currency.ETB,
	region:       RegionEasternAfrica,
	flagEmoji:    "🇪🇹",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataEthiopia) }

var Ethiopia = dataEthiopia
