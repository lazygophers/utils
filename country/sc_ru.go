//go:build (lang_ru || lang_all) && (country_africa || country_all || country_eastern_africa || country_sc)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSeychelles.RegisterName(xlanguage.Russian, "Сейшельские Острова")
	dataSeychelles.RegisterOfficialName(xlanguage.Russian, "Республика Сейшельские Острова")
	dataSeychelles.RegisterCapital(xlanguage.Russian, "Виктория")
}
