//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBelgium.RegisterName(xlanguage.Spanish, "Bélgica")
	dataBelgium.RegisterOfficialName(xlanguage.Spanish, "Reino de Bélgica")
	dataBelgium.RegisterCapital(xlanguage.Spanish, "Bruselas")
}
