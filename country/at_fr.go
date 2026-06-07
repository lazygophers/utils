//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAustria.RegisterName(xlanguage.French, "Autriche")
	dataAustria.RegisterOfficialName(xlanguage.French, "République d'Autriche")
	dataAustria.RegisterCapital(xlanguage.French, "Vienne")
}
