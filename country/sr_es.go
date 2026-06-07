//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSuriname.RegisterName(xlanguage.Spanish, "Surinam")
	dataSuriname.RegisterOfficialName(xlanguage.Spanish, "República de Surinam")
	dataSuriname.RegisterCapital(xlanguage.Spanish, "Paramaribo")
}
