package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Russia — Russian Federation.
var dataRussia = &Country{
	alpha2:       "RU",
	alpha3:       "RUS",
	numeric:      643,
	callingCodes: []string{"+7"},
	timezones:    []string{
		"Europe/Moscow",
		"Europe/Kaliningrad",
		"Europe/Samara",
		"Asia/Yekaterinburg",
		"Asia/Omsk",
		"Asia/Krasnoyarsk",
		"Asia/Irkutsk",
		"Asia/Yakutsk",
		"Asia/Vladivostok",
		"Asia/Magadan",
		"Asia/Kamchatka",
		"Asia/Anadyr",
	},
	tlds:         []string{
		".ru",
		".рф",
	},
	officialLanguage:  xlanguage.Russian,
	spokenLanguages:   []xlanguage.Tag{xlanguage.Russian},
	currency:     currency.RUB,
	region:       RegionEasternEurope,
	flagEmoji:    "🇷🇺",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataRussia) }

var Russia = dataRussia
