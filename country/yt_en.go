package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMayotte.RegisterName(xlanguage.English, "Mayotte")
	dataMayotte.RegisterOfficialName(xlanguage.English, "Department of Mayotte")
	dataMayotte.RegisterCapital(xlanguage.English, "Mamoudzou")
}
