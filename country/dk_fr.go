//go:build (lang_fr || lang_all) && (country_all || country_dk || country_europe || country_northern_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataDenmark.RegisterName(xlanguage.French, "Danemark")
	dataDenmark.RegisterOfficialName(xlanguage.French, "Royaume de Danemark")
	dataDenmark.RegisterCapital(xlanguage.French, "Copenhague")
}
