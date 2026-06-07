//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBulgaria.RegisterName(xlanguage.Spanish, "Bulgaria")
	dataBulgaria.RegisterOfficialName(xlanguage.Spanish, "República de Bulgaria")
	dataBulgaria.RegisterCapital(xlanguage.Spanish, "Sofía")
}
