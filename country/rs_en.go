package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSerbia.RegisterName(xlanguage.English, "Serbia")
	dataSerbia.RegisterOfficialName(xlanguage.English, "Republic of Serbia")
	dataSerbia.RegisterCapital(xlanguage.English, "Belgrade")
}
