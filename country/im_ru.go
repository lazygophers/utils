//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataIsleOfMan.RegisterName(xlanguage.Russian, "Остров Мэн")
	dataIsleOfMan.RegisterOfficialName(xlanguage.Russian, "Остров Мэн")
	dataIsleOfMan.RegisterCapital(xlanguage.Russian, "Дуглас")
}
