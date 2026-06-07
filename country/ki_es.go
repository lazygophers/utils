//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataKiribati.RegisterName(xlanguage.Spanish, "Kiribati")
	dataKiribati.RegisterOfficialName(xlanguage.Spanish, "República de Kiribati")
	dataKiribati.RegisterCapital(xlanguage.Spanish, "Tarawa")
}
