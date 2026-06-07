package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataIndia.RegisterName(xlanguage.English, "India")
	dataIndia.RegisterOfficialName(xlanguage.English, "Republic of India")
	dataIndia.RegisterCapital(xlanguage.English, "New Delhi")
}
