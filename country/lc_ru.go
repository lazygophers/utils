//go:build (lang_ru || lang_all) && (country_all || country_americas || country_caribbean || country_lc)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSaintLucia.RegisterName(xlanguage.Russian, "Сент-Люсия")
	dataSaintLucia.RegisterOfficialName(xlanguage.Russian, "Сент-Люсия")
	dataSaintLucia.RegisterCapital(xlanguage.Russian, "Кастри")
}
