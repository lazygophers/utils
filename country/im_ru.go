//go:build (lang_ru || lang_all) && (country_all || country_europe || country_im || country_northern_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataIsleOfMan.RegisterName(xlanguage.Russian, "Остров Мэн")
	dataIsleOfMan.RegisterOfficialName(xlanguage.Russian, "Остров Мэн")
	dataIsleOfMan.RegisterCapital(xlanguage.Russian, "Дуглас")
}
