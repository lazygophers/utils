//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataIran.RegisterName(xlanguage.French, "Iran")
	dataIran.RegisterOfficialName(xlanguage.French, "République islamique d'Iran")
	dataIran.RegisterCapital(xlanguage.French, "Téhéran")
}
