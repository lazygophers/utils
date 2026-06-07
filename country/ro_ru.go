//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataRomania.RegisterName(xlanguage.Russian, "Румыния")
	dataRomania.RegisterOfficialName(xlanguage.Russian, "Румыния")
	dataRomania.RegisterCapital(xlanguage.Russian, "Бухарест")
}
