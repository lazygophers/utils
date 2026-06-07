//go:build (lang_ru || lang_all) && (country_all || country_eastern_europe || country_europe || country_sk)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSlovakia.RegisterName(xlanguage.Russian, "Словакия")
	dataSlovakia.RegisterOfficialName(xlanguage.Russian, "Словацкая Республика")
	dataSlovakia.RegisterCapital(xlanguage.Russian, "Братислава")
}
