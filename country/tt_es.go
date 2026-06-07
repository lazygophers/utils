//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTrinidadAndTobago.RegisterName(xlanguage.Spanish, "Trinidad y Tobago")
	dataTrinidadAndTobago.RegisterOfficialName(xlanguage.Spanish, "República de Trinidad y Tobago")
	dataTrinidadAndTobago.RegisterCapital(xlanguage.Spanish, "Puerto España")
}
