//go:build country_all || country_americas || country_br || country_south_america

package country

import (
	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/currency"
)

// Brazil — Federative Republic of Brazil.
var dataBrazil = &Country{
	alpha2:       "BR",
	alpha3:       "BRA",
	numeric:      76,
	callingCodes: []string{"+55"},
	timezones:    []string{
		"America/Sao_Paulo",
		"America/Manaus",
		"America/Fortaleza",
		"America/Recife",
		"America/Bahia",
		"America/Belem",
		"America/Campo_Grande",
		"America/Cuiaba",
		"America/Maceio",
		"America/Noronha",
		"America/Porto_Velho",
		"America/Rio_Branco",
		"America/Santarem",
	},
	tlds:         []string{".br"},
	languages:    []xlanguage.Tag{xlanguage.Portuguese},
	currency:     currency.Brl,
	continent:    "SA",
	region:       "Americas",
	subregion:    "South America",
	flagEmoji:    "🇧🇷",
	names:        make(map[xlanguage.Tag]string),
	official:     make(map[xlanguage.Tag]string),
	capital:      make(map[xlanguage.Tag]string),
}

func init() { register(dataBrazil) }

var Brazil = dataBrazil
