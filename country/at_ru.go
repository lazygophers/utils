//go:build (lang_ru || lang_all) && (country_all || country_at || country_europe || country_western_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAustria.RegisterName(xlanguage.Russian, "Австрия")
	dataAustria.RegisterOfficialName(xlanguage.Russian, "Австрийская Республика")
	dataAustria.RegisterCapital(xlanguage.Russian, "Вена")
}
