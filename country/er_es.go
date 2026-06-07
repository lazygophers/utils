//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataEritrea.RegisterName(xlanguage.Spanish, "Eritrea")
	dataEritrea.RegisterOfficialName(xlanguage.Spanish, "Estado de Eritrea")
	dataEritrea.RegisterCapital(xlanguage.Spanish, "Asmara")
}
