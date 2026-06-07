//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataZambia.RegisterName(xlanguage.Russian, "Замбия")
	dataZambia.RegisterOfficialName(xlanguage.Russian, "Республика Замбия")
	dataZambia.RegisterCapital(xlanguage.Russian, "Лусака")
}
