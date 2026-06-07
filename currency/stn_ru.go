//go:build (lang_ru || lang_all) && (country_africa || country_all || country_middle_africa || country_st || currency_all || currency_stn)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Stn.RegisterName(xlanguage.Russian, "Добра")
}
