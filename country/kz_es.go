//go:build (lang_es || lang_all) && (country_all || country_asia || country_central_asia || country_kz)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataKazakhstan.RegisterName(xlanguage.Spanish, "Kazajistán")
	dataKazakhstan.RegisterOfficialName(xlanguage.Spanish, "República de Kazajistán")
	dataKazakhstan.RegisterCapital(xlanguage.Spanish, "Astaná")
}
