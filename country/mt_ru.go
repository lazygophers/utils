//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMalta.RegisterName(xlanguage.Russian, "Мальта")
	dataMalta.RegisterOfficialName(xlanguage.Russian, "Республика Мальта")
	dataMalta.RegisterCapital(xlanguage.Russian, "Валлетта")
}
