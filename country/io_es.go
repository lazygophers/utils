//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBritishIndianOceanTerritory.RegisterName(xlanguage.Spanish, "Territorio Británico del Océano Índico")
	dataBritishIndianOceanTerritory.RegisterOfficialName(xlanguage.Spanish, "Territorio Británico del Océano Índico")
	dataBritishIndianOceanTerritory.RegisterCapital(xlanguage.Spanish, "Diego García")
}
