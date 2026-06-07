//go:build (lang_fr || lang_all) && (country_all || country_asia || country_central_asia || country_kz)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataKazakhstan.RegisterName(xlanguage.French, "Kazakhstan")
	dataKazakhstan.RegisterOfficialName(xlanguage.French, "République du Kazakhstan")
	dataKazakhstan.RegisterCapital(xlanguage.French, "Astana")
}
