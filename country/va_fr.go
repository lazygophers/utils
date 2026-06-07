//go:build (lang_fr || lang_all) && (country_all || country_europe || country_southern_europe || country_va)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataVaticanCity.RegisterName(xlanguage.French, "Vatican")
	dataVaticanCity.RegisterOfficialName(xlanguage.French, "État de la Cité du Vatican")
	dataVaticanCity.RegisterCapital(xlanguage.French, "Vatican")
}
