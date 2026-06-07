package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMontserrat.RegisterName(xlanguage.English, "Montserrat")
	dataMontserrat.RegisterOfficialName(xlanguage.English, "Montserrat")
	dataMontserrat.RegisterCapital(xlanguage.English, "Plymouth")
}
