//go:build (lang_ru || lang_all) && (country_all || country_asia || country_om || country_western_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataOman.RegisterName(xlanguage.Russian, "Оман")
	dataOman.RegisterOfficialName(xlanguage.Russian, "Султанат Оман")
	dataOman.RegisterCapital(xlanguage.Russian, "Маскат")
}
