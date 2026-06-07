package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSanMarino.RegisterName(xlanguage.English, "San Marino")
	dataSanMarino.RegisterOfficialName(xlanguage.English, "Republic of San Marino")
	dataSanMarino.RegisterCapital(xlanguage.English, "San Marino")
}
