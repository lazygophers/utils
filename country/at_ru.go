//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAustria.RegisterName(xlanguage.Russian, "Австрия")
	dataAustria.RegisterOfficialName(xlanguage.Russian, "Австрийская Республика")
	dataAustria.RegisterCapital(xlanguage.Russian, "Вена")
}
