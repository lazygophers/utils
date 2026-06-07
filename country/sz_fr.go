//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataEswatini.RegisterName(xlanguage.French, "Eswatini")
	dataEswatini.RegisterOfficialName(xlanguage.French, "Royaume d'Eswatini")
	dataEswatini.RegisterCapital(xlanguage.French, "Mbabane")
}
