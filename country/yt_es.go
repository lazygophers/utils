//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMayotte.RegisterName(xlanguage.Spanish, "Mayotte")
	dataMayotte.RegisterOfficialName(xlanguage.Spanish, "Departamento de Mayotte")
	dataMayotte.RegisterCapital(xlanguage.Spanish, "Mamoudzou")
}
