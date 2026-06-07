//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSierraLeone.RegisterName(xlanguage.Spanish, "Sierra Leona")
	dataSierraLeone.RegisterOfficialName(xlanguage.Spanish, "República de Sierra Leona")
	dataSierraLeone.RegisterCapital(xlanguage.Spanish, "Freetown")
}
