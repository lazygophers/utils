//go:build (lang_fr || lang_all) && (country_all || country_americas || country_co || country_south_america)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataColombia.RegisterName(xlanguage.French, "Colombie")
	dataColombia.RegisterOfficialName(xlanguage.French, "République de Colombie")
	dataColombia.RegisterCapital(xlanguage.French, "Bogota")
}
