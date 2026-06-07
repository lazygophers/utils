package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGuernsey.RegisterName(xlanguage.English, "Guernsey")
	dataGuernsey.RegisterOfficialName(xlanguage.English, "Bailiwick of Guernsey")
	dataGuernsey.RegisterCapital(xlanguage.English, "Saint Peter Port")
}
