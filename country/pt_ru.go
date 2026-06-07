//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataPortugal.RegisterName(xlanguage.Russian, "Португалия")
	dataPortugal.RegisterOfficialName(xlanguage.Russian, "Португальская Республика")
	dataPortugal.RegisterCapital(xlanguage.Russian, "Лиссабон")
}
