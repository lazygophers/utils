//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSouthKorea.RegisterName(xlanguage.Spanish, "Corea del Sur")
	dataSouthKorea.RegisterOfficialName(xlanguage.Spanish, "República de Corea")
	dataSouthKorea.RegisterCapital(xlanguage.Spanish, "Seúl")
}
