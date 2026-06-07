//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataLebanon.RegisterName(xlanguage.Spanish, "Líbano")
	dataLebanon.RegisterOfficialName(xlanguage.Spanish, "República Libanesa")
	dataLebanon.RegisterCapital(xlanguage.Spanish, "Beirut")
}
