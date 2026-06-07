//go:build (lang_ru || lang_all) && (country_all || country_asia || country_il || country_western_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataIsrael.RegisterName(xlanguage.Russian, "Израиль")
	dataIsrael.RegisterOfficialName(xlanguage.Russian, "Государство Израиль")
	dataIsrael.RegisterCapital(xlanguage.Russian, "Иерусалим")
}
