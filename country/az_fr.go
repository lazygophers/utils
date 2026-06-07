//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAzerbaijan.RegisterName(xlanguage.French, "Azerbaïdjan")
	dataAzerbaijan.RegisterOfficialName(xlanguage.French, "République d'Azerbaïdjan")
	dataAzerbaijan.RegisterCapital(xlanguage.French, "Bakou")
}
