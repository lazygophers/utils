//go:build (lang_fr || lang_all) && (country_all || country_cz || country_eastern_europe || country_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCzechia.RegisterName(xlanguage.French, "Tchéquie")
	dataCzechia.RegisterOfficialName(xlanguage.French, "République tchèque")
	dataCzechia.RegisterCapital(xlanguage.French, "Prague")
}
