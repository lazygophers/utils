//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataIceland.RegisterName(xlanguage.Spanish, "Islandia")
	dataIceland.RegisterOfficialName(xlanguage.Spanish, "Islandia")
	dataIceland.RegisterCapital(xlanguage.Spanish, "Reikiavik")
}
