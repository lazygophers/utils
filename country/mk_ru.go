//go:build (lang_ru || lang_all) && (country_all || country_europe || country_mk || country_southern_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNorthMacedonia.RegisterName(xlanguage.Russian, "Северная Македония")
	dataNorthMacedonia.RegisterOfficialName(xlanguage.Russian, "Республика Северная Македония")
	dataNorthMacedonia.RegisterCapital(xlanguage.Russian, "Скопье")
}
