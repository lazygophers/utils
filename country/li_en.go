package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataLiechtenstein.RegisterName(xlanguage.English, "Liechtenstein")
	dataLiechtenstein.RegisterOfficialName(xlanguage.English, "Principality of Liechtenstein")
	dataLiechtenstein.RegisterCapital(xlanguage.English, "Vaduz")
}
