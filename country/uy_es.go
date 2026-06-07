package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataUruguay.RegisterName(xlanguage.Spanish, "Uruguay")
	dataUruguay.RegisterOfficialName(xlanguage.Spanish, "República Oriental del Uruguay")
	dataUruguay.RegisterCapital(xlanguage.Spanish, "Montevideo")
}
