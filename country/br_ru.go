//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBrazil.RegisterName(xlanguage.Russian, "Бразилия")
	dataBrazil.RegisterOfficialName(xlanguage.Russian, "Федеративная Республика Бразилия")
	dataBrazil.RegisterCapital(xlanguage.Russian, "Бразилиа")
}
