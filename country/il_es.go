//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataIsrael.RegisterName(xlanguage.Spanish, "Israel")
	dataIsrael.RegisterOfficialName(xlanguage.Spanish, "Estado de Israel")
	dataIsrael.RegisterCapital(xlanguage.Spanish, "Jerusalén")
}
