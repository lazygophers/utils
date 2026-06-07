//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataFrenchGuiana.RegisterName(xlanguage.Russian, "Французская Гвиана")
	dataFrenchGuiana.RegisterOfficialName(xlanguage.Russian, "Гвиана")
	dataFrenchGuiana.RegisterCapital(xlanguage.Russian, "Кайенна")
}
