//go:build (lang_ru || lang_all) && (country_all || country_cz || country_eastern_europe || country_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCzechia.RegisterName(xlanguage.Russian, "Чехия")
	dataCzechia.RegisterOfficialName(xlanguage.Russian, "Чешская Республика")
	dataCzechia.RegisterCapital(xlanguage.Russian, "Прага")
}
