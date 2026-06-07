//go:build (lang_ru || lang_all) && (country_all || country_asia || country_la || country_south_eastern_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataLaos.RegisterName(xlanguage.Russian, "Лаос")
	dataLaos.RegisterOfficialName(xlanguage.Russian, "Лаосская Народно-Демократическая Республика")
	dataLaos.RegisterCapital(xlanguage.Russian, "Вьентьян")
}
