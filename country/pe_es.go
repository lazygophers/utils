package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataPeru.RegisterName(xlanguage.Spanish, "Perú")
	dataPeru.RegisterOfficialName(xlanguage.Spanish, "República del Perú")
	dataPeru.RegisterCapital(xlanguage.Spanish, "Lima")
}
