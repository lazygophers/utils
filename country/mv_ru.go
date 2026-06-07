//go:build (lang_ru || lang_all) && (country_all || country_asia || country_mv || country_southern_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMaldives.RegisterName(xlanguage.Russian, "Мальдивы")
	dataMaldives.RegisterOfficialName(xlanguage.Russian, "Мальдивская Республика")
	dataMaldives.RegisterCapital(xlanguage.Russian, "Мале")
}
