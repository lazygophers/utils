//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataUganda.RegisterName(xlanguage.Spanish, "Uganda")
	dataUganda.RegisterOfficialName(xlanguage.Spanish, "República de Uganda")
	dataUganda.RegisterCapital(xlanguage.Spanish, "Kampala")
}
