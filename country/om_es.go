//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataOman.RegisterName(xlanguage.Spanish, "Omán")
	dataOman.RegisterOfficialName(xlanguage.Spanish, "Sultanato de Omán")
	dataOman.RegisterCapital(xlanguage.Spanish, "Mascate")
}
