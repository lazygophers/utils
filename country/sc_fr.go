package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSeychelles.RegisterName(xlanguage.French, "Seychelles")
	dataSeychelles.RegisterOfficialName(xlanguage.French, "République des Seychelles")
	dataSeychelles.RegisterCapital(xlanguage.French, "Victoria")
}
