//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBenin.RegisterName(xlanguage.Spanish, "Benín")
	dataBenin.RegisterOfficialName(xlanguage.Spanish, "República de Benín")
	dataBenin.RegisterCapital(xlanguage.Spanish, "Porto Novo")
}
