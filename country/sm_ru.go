//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSanMarino.RegisterName(xlanguage.Russian, "Сан-Марино")
	dataSanMarino.RegisterOfficialName(xlanguage.Russian, "Республика Сан-Марино")
	dataSanMarino.RegisterCapital(xlanguage.Russian, "Сан-Марино")
}
