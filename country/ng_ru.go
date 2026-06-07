//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNigeria.RegisterName(xlanguage.Russian, "Нигерия")
	dataNigeria.RegisterOfficialName(xlanguage.Russian, "Федеративная Республика Нигерия")
	dataNigeria.RegisterCapital(xlanguage.Russian, "Абуджа")
}
