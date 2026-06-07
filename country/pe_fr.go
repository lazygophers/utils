//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataPeru.RegisterName(xlanguage.French, "Pérou")
	dataPeru.RegisterOfficialName(xlanguage.French, "République du Pérou")
	dataPeru.RegisterCapital(xlanguage.French, "Lima")
}
