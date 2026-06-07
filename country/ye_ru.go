//go:build (lang_ru || lang_all) && (country_all || country_asia || country_western_asia || country_ye)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataYemen.RegisterName(xlanguage.Russian, "Йемен")
	dataYemen.RegisterOfficialName(xlanguage.Russian, "Йеменская Республика")
	dataYemen.RegisterCapital(xlanguage.Russian, "Сана")
}
