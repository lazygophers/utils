//go:build (lang_fr || lang_all) && (country_all || country_australia_and_new_zealand || country_nf || country_oceania)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNorfolkIsland.RegisterName(xlanguage.French, "Île Norfolk")
	dataNorfolkIsland.RegisterOfficialName(xlanguage.French, "Territoire de l'île Norfolk")
	dataNorfolkIsland.RegisterCapital(xlanguage.French, "Kingston")
}
