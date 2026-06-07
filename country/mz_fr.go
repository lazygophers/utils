//go:build (lang_fr || lang_all) && (country_africa || country_all || country_eastern_africa || country_mz)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMozambique.RegisterName(xlanguage.French, "Mozambique")
	dataMozambique.RegisterOfficialName(xlanguage.French, "République du Mozambique")
	dataMozambique.RegisterCapital(xlanguage.French, "Maputo")
}
