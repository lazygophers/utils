//go:build (lang_es || lang_all) && (country_africa || country_all || country_ly || country_northern_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataLibya.RegisterName(xlanguage.Spanish, "Libia")
	dataLibya.RegisterOfficialName(xlanguage.Spanish, "Estado de Libia")
	dataLibya.RegisterCapital(xlanguage.Spanish, "Trípoli")
}
