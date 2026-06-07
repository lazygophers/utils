//go:build (lang_ru || lang_all) && (country_all || country_australia_and_new_zealand || country_nf || country_oceania)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNorfolkIsland.RegisterName(xlanguage.Russian, "Остров Норфолк")
	dataNorfolkIsland.RegisterOfficialName(xlanguage.Russian, "Территория Остров Норфолк")
	dataNorfolkIsland.RegisterCapital(xlanguage.Russian, "Кингстон")
}
