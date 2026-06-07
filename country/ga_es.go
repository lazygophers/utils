//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGabon.RegisterName(xlanguage.Spanish, "Gabón")
	dataGabon.RegisterOfficialName(xlanguage.Spanish, "República Gabonesa")
	dataGabon.RegisterCapital(xlanguage.Spanish, "Libreville")
}
