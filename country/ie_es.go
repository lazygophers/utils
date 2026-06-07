//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataIreland.RegisterName(xlanguage.Spanish, "Irlanda")
	dataIreland.RegisterOfficialName(xlanguage.Spanish, "Irlanda")
	dataIreland.RegisterCapital(xlanguage.Spanish, "Dublín")
}
