//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBotswana.RegisterName(xlanguage.Russian, "Ботсвана")
	dataBotswana.RegisterOfficialName(xlanguage.Russian, "Республика Ботсвана")
	dataBotswana.RegisterCapital(xlanguage.Russian, "Габороне")
}
