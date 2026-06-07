//go:build (lang_fr || lang_all) && (country_all || country_europe || country_rs || country_southern_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSerbia.RegisterName(xlanguage.French, "Serbie")
	dataSerbia.RegisterOfficialName(xlanguage.French, "République de Serbie")
	dataSerbia.RegisterCapital(xlanguage.French, "Belgrade")
}
