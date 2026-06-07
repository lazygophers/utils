//go:build (lang_es || lang_all) && (country_africa || country_all || country_eastern_africa || country_mz)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMozambique.RegisterName(xlanguage.Spanish, "Mozambique")
	dataMozambique.RegisterOfficialName(xlanguage.Spanish, "República de Mozambique")
	dataMozambique.RegisterCapital(xlanguage.Spanish, "Maputo")
}
