package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGreece.RegisterName(xlanguage.English, "Greece")
	dataGreece.RegisterOfficialName(xlanguage.English, "Hellenic Republic")
	dataGreece.RegisterCapital(xlanguage.English, "Athens")
}
