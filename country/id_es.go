//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataIndonesia.RegisterName(xlanguage.Spanish, "Indonesia")
	dataIndonesia.RegisterOfficialName(xlanguage.Spanish, "República de Indonesia")
	dataIndonesia.RegisterCapital(xlanguage.Spanish, "Yakarta")
}
