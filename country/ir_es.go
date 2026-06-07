//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataIran.RegisterName(xlanguage.Spanish, "Irán")
	dataIran.RegisterOfficialName(xlanguage.Spanish, "República Islámica de Irán")
	dataIran.RegisterCapital(xlanguage.Spanish, "Teherán")
}
