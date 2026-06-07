//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMadagascar.RegisterName(xlanguage.Spanish, "Madagascar")
	dataMadagascar.RegisterOfficialName(xlanguage.Spanish, "República de Madagascar")
	dataMadagascar.RegisterCapital(xlanguage.Spanish, "Antananarivo")
}
