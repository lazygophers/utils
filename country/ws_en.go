package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSamoa.RegisterName(xlanguage.English, "Samoa")
	dataSamoa.RegisterOfficialName(xlanguage.English, "Independent State of Samoa")
	dataSamoa.RegisterCapital(xlanguage.English, "Apia")
}
