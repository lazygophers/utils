//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataPalau.RegisterName(xlanguage.Spanish, "Palaos")
	dataPalau.RegisterOfficialName(xlanguage.Spanish, "República de Palaos")
	dataPalau.RegisterCapital(xlanguage.Spanish, "Ngerulmud")
}
