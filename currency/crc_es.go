//go:build (lang_es || lang_all) && (country_all || country_americas || country_central_america || country_cr || currency_all || currency_crc)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Crc.RegisterName(xlanguage.Spanish, "Colón costarricense")
}
