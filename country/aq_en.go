package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAntarctica.RegisterName(xlanguage.English, "Antarctica")
	dataAntarctica.RegisterOfficialName(xlanguage.English, "Antarctica")
}
