//go:build (lang_ru || lang_all) && (country_all || country_europe || country_lv || country_northern_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataLatvia.RegisterName(xlanguage.Russian, "Латвия")
	dataLatvia.RegisterOfficialName(xlanguage.Russian, "Латвийская Республика")
	dataLatvia.RegisterCapital(xlanguage.Russian, "Рига")
}
