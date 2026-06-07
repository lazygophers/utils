//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMauritania.RegisterName(xlanguage.Spanish, "Mauritania")
	dataMauritania.RegisterOfficialName(xlanguage.Spanish, "República Islámica de Mauritania")
	dataMauritania.RegisterCapital(xlanguage.Spanish, "Nuakchot")
}
