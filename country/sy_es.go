//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSyria.RegisterName(xlanguage.Spanish, "Siria")
	dataSyria.RegisterOfficialName(xlanguage.Spanish, "República Árabe Siria")
	dataSyria.RegisterCapital(xlanguage.Spanish, "Damasco")
}
