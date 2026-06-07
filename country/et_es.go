//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataEthiopia.RegisterName(xlanguage.Spanish, "Etiopía")
	dataEthiopia.RegisterOfficialName(xlanguage.Spanish, "República Federal Democrática de Etiopía")
	dataEthiopia.RegisterCapital(xlanguage.Spanish, "Adís Abeba")
}
