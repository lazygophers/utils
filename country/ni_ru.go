//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNicaragua.RegisterName(xlanguage.Russian, "Никарагуа")
	dataNicaragua.RegisterOfficialName(xlanguage.Russian, "Республика Никарагуа")
	dataNicaragua.RegisterCapital(xlanguage.Russian, "Манагуа")
}
