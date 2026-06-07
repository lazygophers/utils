//go:build (lang_ru || lang_all) && (country_all || country_eastern_europe || country_europe || country_pl)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataPoland.RegisterName(xlanguage.Russian, "Польша")
	dataPoland.RegisterOfficialName(xlanguage.Russian, "Республика Польша")
	dataPoland.RegisterCapital(xlanguage.Russian, "Варшава")
}
