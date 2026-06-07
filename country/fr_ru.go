//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataFrance.RegisterName(xlanguage.Russian, "Франция")
	dataFrance.RegisterOfficialName(xlanguage.Russian, "Французская Республика")
	dataFrance.RegisterCapital(xlanguage.Russian, "Париж")
}
