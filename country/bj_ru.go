//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBenin.RegisterName(xlanguage.Russian, "Бенин")
	dataBenin.RegisterOfficialName(xlanguage.Russian, "Республика Бенин")
	dataBenin.RegisterCapital(xlanguage.Russian, "Порто-Ново")
}
