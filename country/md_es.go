//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMoldova.RegisterName(xlanguage.Spanish, "Moldavia")
	dataMoldova.RegisterOfficialName(xlanguage.Spanish, "República de Moldavia")
	dataMoldova.RegisterCapital(xlanguage.Spanish, "Chisináu")
}
