//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataDominica.RegisterName(xlanguage.French, "Dominique")
	dataDominica.RegisterOfficialName(xlanguage.French, "Commonwealth de Dominique")
	dataDominica.RegisterCapital(xlanguage.French, "Roseau")
}
