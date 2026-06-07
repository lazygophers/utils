package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNorfolkIsland.RegisterName(xlanguage.English, "Norfolk Island")
	dataNorfolkIsland.RegisterOfficialName(xlanguage.English, "Territory of Norfolk Island")
	dataNorfolkIsland.RegisterCapital(xlanguage.English, "Kingston")
}
