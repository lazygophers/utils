//go:build (lang_es || lang_all) && (country_all || country_ch || country_europe || country_western_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSwitzerland.RegisterName(xlanguage.Spanish, "Suiza")
	dataSwitzerland.RegisterOfficialName(xlanguage.Spanish, "Confederación Suiza")
	dataSwitzerland.RegisterCapital(xlanguage.Spanish, "Berna")
}
