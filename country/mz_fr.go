//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMozambique.RegisterName(xlanguage.French, "Mozambique")
	dataMozambique.RegisterOfficialName(xlanguage.French, "République du Mozambique")
	dataMozambique.RegisterCapital(xlanguage.French, "Maputo")
}
