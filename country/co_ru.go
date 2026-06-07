//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataColombia.RegisterName(xlanguage.Russian, "Колумбия")
	dataColombia.RegisterOfficialName(xlanguage.Russian, "Республика Колумбия")
	dataColombia.RegisterCapital(xlanguage.Russian, "Богота")
}
