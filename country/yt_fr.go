package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMayotte.RegisterName(xlanguage.French, "Mayotte")
	dataMayotte.RegisterOfficialName(xlanguage.French, "Département de Mayotte")
	dataMayotte.RegisterCapital(xlanguage.French, "Mamoudzou")
}
