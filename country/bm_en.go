package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBermuda.RegisterName(xlanguage.English, "Bermuda")
	dataBermuda.RegisterOfficialName(xlanguage.English, "Bermuda")
	dataBermuda.RegisterCapital(xlanguage.English, "Hamilton")
}
