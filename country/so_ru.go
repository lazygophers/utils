//go:build (lang_ru || lang_all) && (country_africa || country_all || country_eastern_africa || country_so)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSomalia.RegisterName(xlanguage.Russian, "Сомали")
	dataSomalia.RegisterOfficialName(xlanguage.Russian, "Федеративная Республика Сомали")
	dataSomalia.RegisterCapital(xlanguage.Russian, "Могадишо")
}
