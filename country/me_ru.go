//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMontenegro.RegisterName(xlanguage.Russian, "Черногория")
	dataMontenegro.RegisterOfficialName(xlanguage.Russian, "Черногория")
	dataMontenegro.RegisterCapital(xlanguage.Russian, "Подгорица")
}
