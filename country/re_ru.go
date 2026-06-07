//go:build (lang_ru || lang_all) && (country_africa || country_all || country_eastern_africa || country_re)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataReunion.RegisterName(xlanguage.Russian, "Реюньон")
	dataReunion.RegisterOfficialName(xlanguage.Russian, "Реюньон")
	dataReunion.RegisterCapital(xlanguage.Russian, "Сен-Дени")
}
