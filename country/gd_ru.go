//go:build (lang_ru || lang_all) && (country_all || country_americas || country_caribbean || country_gd)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGrenada.RegisterName(xlanguage.Russian, "Гренада")
	dataGrenada.RegisterOfficialName(xlanguage.Russian, "Гренада")
	dataGrenada.RegisterCapital(xlanguage.Russian, "Сент-Джорджес")
}
