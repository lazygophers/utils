//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataJapan.RegisterName(xlanguage.Russian, "Япония")
	dataJapan.RegisterOfficialName(xlanguage.Russian, "Япония")
	dataJapan.RegisterCapital(xlanguage.Russian, "Токио")
}
