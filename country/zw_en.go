package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataZimbabwe.RegisterName(xlanguage.English, "Zimbabwe")
	dataZimbabwe.RegisterOfficialName(xlanguage.English, "Republic of Zimbabwe")
	dataZimbabwe.RegisterCapital(xlanguage.English, "Harare")
}
