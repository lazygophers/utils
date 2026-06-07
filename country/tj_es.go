//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTajikistan.RegisterName(xlanguage.Spanish, "Tayikistán")
	dataTajikistan.RegisterOfficialName(xlanguage.Spanish, "República de Tayikistán")
	dataTajikistan.RegisterCapital(xlanguage.Spanish, "Dusambé")
}
