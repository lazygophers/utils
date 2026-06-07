//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCentralAfricanRepublic.RegisterName(xlanguage.Spanish, "República Centroafricana")
	dataCentralAfricanRepublic.RegisterOfficialName(xlanguage.Spanish, "República Centroafricana")
	dataCentralAfricanRepublic.RegisterCapital(xlanguage.Spanish, "Bangui")
}
