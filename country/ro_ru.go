//go:build (lang_ru || lang_all) && (country_all || country_eastern_europe || country_europe || country_ro)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataRomania.RegisterName(xlanguage.Russian, "Румыния")
	dataRomania.RegisterOfficialName(xlanguage.Russian, "Румыния")
	dataRomania.RegisterCapital(xlanguage.Russian, "Бухарест")
}
