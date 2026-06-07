//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataDominicanRepublic.RegisterName(xlanguage.Russian, "Доминиканская Республика")
	dataDominicanRepublic.RegisterOfficialName(xlanguage.Russian, "Доминиканская Республика")
	dataDominicanRepublic.RegisterCapital(xlanguage.Russian, "Санто-Доминго")
}
