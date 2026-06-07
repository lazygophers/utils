//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataJamaica.RegisterName(xlanguage.French, "Jamaïque")
	dataJamaica.RegisterOfficialName(xlanguage.French, "Jamaïque")
	dataJamaica.RegisterCapital(xlanguage.French, "Kingston")
}
