//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCaboVerde.RegisterName(xlanguage.French, "Cap-Vert")
	dataCaboVerde.RegisterOfficialName(xlanguage.French, "République du Cap-Vert")
	dataCaboVerde.RegisterCapital(xlanguage.French, "Praia")
}
