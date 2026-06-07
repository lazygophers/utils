//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTurkey.RegisterName(xlanguage.Russian, "Турция")
	dataTurkey.RegisterOfficialName(xlanguage.Russian, "Турецкая Республика")
	dataTurkey.RegisterCapital(xlanguage.Russian, "Анкара")
}
