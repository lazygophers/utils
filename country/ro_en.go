package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataRomania.RegisterName(xlanguage.English, "Romania")
	dataRomania.RegisterOfficialName(xlanguage.English, "Romania")
	dataRomania.RegisterCapital(xlanguage.English, "Bucharest")
}
