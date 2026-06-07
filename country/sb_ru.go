//go:build (lang_ru || lang_all) && (country_all || country_melanesia || country_oceania || country_sb)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSolomonIslands.RegisterName(xlanguage.Russian, "Соломоновы Острова")
	dataSolomonIslands.RegisterOfficialName(xlanguage.Russian, "Соломоновы Острова")
	dataSolomonIslands.RegisterCapital(xlanguage.Russian, "Хониара")
}
