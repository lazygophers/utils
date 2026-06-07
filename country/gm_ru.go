//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGambia.RegisterName(xlanguage.Russian, "Гамбия")
	dataGambia.RegisterOfficialName(xlanguage.Russian, "Республика Гамбия")
	dataGambia.RegisterCapital(xlanguage.Russian, "Банжул")
}
