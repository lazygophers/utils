//go:build (lang_ru || lang_all) && (country_all || country_americas || country_caribbean || country_jm)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataJamaica.RegisterName(xlanguage.Russian, "Ямайка")
	dataJamaica.RegisterOfficialName(xlanguage.Russian, "Ямайка")
	dataJamaica.RegisterCapital(xlanguage.Russian, "Кингстон")
}
