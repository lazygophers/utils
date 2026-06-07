//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCyprus.RegisterName(xlanguage.Russian, "Кипр")
	dataCyprus.RegisterOfficialName(xlanguage.Russian, "Республика Кипр")
	dataCyprus.RegisterCapital(xlanguage.Russian, "Никосия")
}
