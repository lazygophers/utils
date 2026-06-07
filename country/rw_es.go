//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataRwanda.RegisterName(xlanguage.Spanish, "Ruanda")
	dataRwanda.RegisterOfficialName(xlanguage.Spanish, "República de Ruanda")
	dataRwanda.RegisterCapital(xlanguage.Spanish, "Kigali")
}
