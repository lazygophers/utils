//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataFiji.RegisterName(xlanguage.Spanish, "Fiyi")
	dataFiji.RegisterOfficialName(xlanguage.Spanish, "República de Fiyi")
	dataFiji.RegisterCapital(xlanguage.Spanish, "Suva")
}
