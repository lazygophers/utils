//go:build (lang_ru || lang_all) && (country_africa || country_all || country_eastern_africa || country_sc || currency_all || currency_scr)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	SCR.RegisterName(xlanguage.Russian, "Сейшельская рупия")
}
