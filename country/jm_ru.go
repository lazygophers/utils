//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataJamaica.RegisterName(xlanguage.Russian, "Ямайка")
	dataJamaica.RegisterOfficialName(xlanguage.Russian, "Ямайка")
	dataJamaica.RegisterCapital(xlanguage.Russian, "Кингстон")
}
