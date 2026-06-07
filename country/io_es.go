//go:build (lang_es || lang_all) && (country_all || country_asia || country_eastern_africa || country_io)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBritishIndianOceanTerritory.RegisterName(xlanguage.Spanish, "Territorio Británico del Océano Índico")
	dataBritishIndianOceanTerritory.RegisterOfficialName(xlanguage.Spanish, "Territorio Británico del Océano Índico")
	dataBritishIndianOceanTerritory.RegisterCapital(xlanguage.Spanish, "Diego García")
}
