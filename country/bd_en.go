package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBangladesh.RegisterName(xlanguage.English, "Bangladesh")
	dataBangladesh.RegisterOfficialName(xlanguage.English, "People's Republic of Bangladesh")
	dataBangladesh.RegisterCapital(xlanguage.English, "Dhaka")
}
