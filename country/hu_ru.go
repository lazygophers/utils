//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataHungary.RegisterName(xlanguage.Russian, "Венгрия")
	dataHungary.RegisterOfficialName(xlanguage.Russian, "Венгрия")
	dataHungary.RegisterCapital(xlanguage.Russian, "Будапешт")
}
