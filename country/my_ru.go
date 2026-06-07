//go:build (lang_ru || lang_all) && (country_all || country_asia || country_my || country_south_eastern_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMalaysia.RegisterName(xlanguage.Russian, "Малайзия")
	dataMalaysia.RegisterOfficialName(xlanguage.Russian, "Малайзия")
	dataMalaysia.RegisterCapital(xlanguage.Russian, "Куала-Лумпур")
}
