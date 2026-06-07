//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataHonduras.RegisterName(xlanguage.French, "Honduras")
	dataHonduras.RegisterOfficialName(xlanguage.French, "République du Honduras")
	dataHonduras.RegisterCapital(xlanguage.French, "Tegucigalpa")
}
