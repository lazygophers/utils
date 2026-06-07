//go:build country_all || country_asia || country_south_eastern_asia || country_vn

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Vietnam — Socialist Republic of Viet Nam.
var dataVietnam = &Country{
	alpha2:       "VN",
	alpha3:       "VNM",
	numeric:      704,
	callingCodes: []string{"+84"},
	timezones:    []string{"Asia/Ho_Chi_Minh"},
	tlds:         []string{".vn"},
	officialLanguage:  xlanguage.Vietnamese,
	spokenLanguages:   []xlanguage.Tag{xlanguage.Vietnamese},
	currency:     currency.VND,
	region:       RegionSouthEasternAsia,
	flagEmoji:    "🇻🇳",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataVietnam) }

var Vietnam = dataVietnam
