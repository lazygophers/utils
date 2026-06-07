//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataIvoryCoast.RegisterName(xlanguage.Spanish, "Costa de Marfil")
	dataIvoryCoast.RegisterOfficialName(xlanguage.Spanish, "República de Costa de Marfil")
	dataIvoryCoast.RegisterCapital(xlanguage.Spanish, "Yamusukro")
}
