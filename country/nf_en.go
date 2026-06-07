//go:build country_all || country_australia_and_new_zealand || country_nf || country_oceania

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNorfolkIsland.RegisterName(xlanguage.English, "Norfolk Island")
	dataNorfolkIsland.RegisterOfficialName(xlanguage.English, "Territory of Norfolk Island")
	dataNorfolkIsland.RegisterCapital(xlanguage.English, "Kingston")
}
