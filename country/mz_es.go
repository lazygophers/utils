//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMozambique.RegisterName(xlanguage.Spanish, "Mozambique")
	dataMozambique.RegisterOfficialName(xlanguage.Spanish, "República de Mozambique")
	dataMozambique.RegisterCapital(xlanguage.Spanish, "Maputo")
}
