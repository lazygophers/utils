//go:build (lang_ru || lang_all) && (country_all || country_oceania || country_polynesia || country_ws)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSamoa.RegisterName(xlanguage.Russian, "Самоа")
	dataSamoa.RegisterOfficialName(xlanguage.Russian, "Независимое Государство Самоа")
	dataSamoa.RegisterCapital(xlanguage.Russian, "Апиа")
}
