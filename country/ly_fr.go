//go:build (lang_fr || lang_all) && (country_africa || country_all || country_ly || country_northern_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataLibya.RegisterName(xlanguage.French, "Libye")
	dataLibya.RegisterOfficialName(xlanguage.French, "État de Libye")
	dataLibya.RegisterCapital(xlanguage.French, "Tripoli")
}
