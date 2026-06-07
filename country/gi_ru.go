//go:build (lang_ru || lang_all) && (country_all || country_europe || country_gi || country_southern_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGibraltar.RegisterName(xlanguage.Russian, "Гибралтар")
	dataGibraltar.RegisterOfficialName(xlanguage.Russian, "Гибралтар")
	dataGibraltar.RegisterCapital(xlanguage.Russian, "Гибралтар")
}
