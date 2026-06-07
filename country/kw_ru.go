//go:build (lang_ru || lang_all) && (country_all || country_asia || country_kw || country_western_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataKuwait.RegisterName(xlanguage.Russian, "Кувейт")
	dataKuwait.RegisterOfficialName(xlanguage.Russian, "Государство Кувейт")
	dataKuwait.RegisterCapital(xlanguage.Russian, "Эль-Кувейт")
}
