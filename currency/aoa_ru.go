//go:build (lang_ru || lang_all) && (country_africa || country_all || country_ao || country_middle_africa || currency_all || currency_aoa)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Aoa.RegisterName(xlanguage.Russian, "Кванза")
}
