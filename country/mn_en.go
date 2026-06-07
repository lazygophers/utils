package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMongolia.RegisterName(xlanguage.English, "Mongolia")
	dataMongolia.RegisterOfficialName(xlanguage.English, "Mongolia")
	dataMongolia.RegisterCapital(xlanguage.English, "Ulaanbaatar")
}
