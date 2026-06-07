//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBritishIndianOceanTerritory.RegisterName(xlanguage.French, "Territoire britannique de l'océan Indien")
	dataBritishIndianOceanTerritory.RegisterOfficialName(xlanguage.French, "Territoire britannique de l'océan Indien")
	dataBritishIndianOceanTerritory.RegisterCapital(xlanguage.French, "Diego Garcia")
}
