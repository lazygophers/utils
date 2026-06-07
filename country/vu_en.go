package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataVanuatu.RegisterName(xlanguage.English, "Vanuatu")
	dataVanuatu.RegisterOfficialName(xlanguage.English, "Republic of Vanuatu")
	dataVanuatu.RegisterCapital(xlanguage.English, "Port Vila")
}
