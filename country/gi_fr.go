//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGibraltar.RegisterName(xlanguage.French, "Gibraltar")
	dataGibraltar.RegisterOfficialName(xlanguage.French, "Gibraltar")
	dataGibraltar.RegisterCapital(xlanguage.French, "Gibraltar")
}
