package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataChristmasIsland.RegisterName(xlanguage.English, "Christmas Island")
	dataChristmasIsland.RegisterOfficialName(xlanguage.English, "Territory of Christmas Island")
	dataChristmasIsland.RegisterCapital(xlanguage.English, "Flying Fish Cove")
}
