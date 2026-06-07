//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTogo.RegisterName(xlanguage.Russian, "Того")
	dataTogo.RegisterOfficialName(xlanguage.Russian, "Тоголезская Республика")
	dataTogo.RegisterCapital(xlanguage.Russian, "Ломе")
}
