//go:build (lang_ru || lang_all) && (country_africa || country_all || country_gw || country_western_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGuineaBissau.RegisterName(xlanguage.Russian, "Гвинея-Бисау")
	dataGuineaBissau.RegisterOfficialName(xlanguage.Russian, "Республика Гвинея-Бисау")
	dataGuineaBissau.RegisterCapital(xlanguage.Russian, "Бисау")
}
