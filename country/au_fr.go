//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAustralia.RegisterName(xlanguage.French, "Australie")
	dataAustralia.RegisterOfficialName(xlanguage.French, "Commonwealth d'Australie")
	dataAustralia.RegisterCapital(xlanguage.French, "Canberra")
}
