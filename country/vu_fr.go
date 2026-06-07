package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataVanuatu.RegisterName(xlanguage.French, "Vanuatu")
	dataVanuatu.RegisterOfficialName(xlanguage.French, "République du Vanuatu")
	dataVanuatu.RegisterCapital(xlanguage.French, "Port-Vila")
}
