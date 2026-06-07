//go:build (lang_ru || lang_all) && (country_all || country_americas || country_caribbean || country_ky || currency_all || currency_kyd)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Kyd.RegisterName(xlanguage.Russian, "Доллар Каймановых Островов")
}
