//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSerbia.RegisterName(xlanguage.Spanish, "Serbia")
	dataSerbia.RegisterOfficialName(xlanguage.Spanish, "República de Serbia")
	dataSerbia.RegisterCapital(xlanguage.Spanish, "Belgrado")
}
