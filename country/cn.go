package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// China — People's Republic of China.
var dataChina = &Country{
	alpha2:       "CN",
	alpha3:       "CHN",
	numeric:      156,
	callingCodes: []string{"+86"},
	timezones:    []string{"Asia/Shanghai"},
	tlds:         []string{".cn", ".中国", ".中國", ".公司", ".网络"},
	languages:    []xlanguage.Tag{xlanguage.Chinese},
	currency:     currency.Cny,
	continent:    "AS",
	region:       "Asia",
	subregion:    "Eastern Asia",
	flagEmoji:    "🇨🇳",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataChina) }
