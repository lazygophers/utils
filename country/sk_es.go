//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSlovakia.RegisterName(xlanguage.Spanish, "Eslovaquia")
	dataSlovakia.RegisterOfficialName(xlanguage.Spanish, "República Eslovaca")
	dataSlovakia.RegisterCapital(xlanguage.Spanish, "Bratislava")
}
