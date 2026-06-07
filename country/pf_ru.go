//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataFrenchPolynesia.RegisterName(xlanguage.Russian, "Французская Полинезия")
	dataFrenchPolynesia.RegisterOfficialName(xlanguage.Russian, "Французская Полинезия")
	dataFrenchPolynesia.RegisterCapital(xlanguage.Russian, "Папеэте")
}
