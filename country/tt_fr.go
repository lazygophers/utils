//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTrinidadAndTobago.RegisterName(xlanguage.French, "Trinité-et-Tobago")
	dataTrinidadAndTobago.RegisterOfficialName(xlanguage.French, "République de Trinité-et-Tobago")
	dataTrinidadAndTobago.RegisterCapital(xlanguage.French, "Port-d'Espagne")
}
