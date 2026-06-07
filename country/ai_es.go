//go:build (lang_es || lang_all) && (country_ai || country_all || country_americas || country_caribbean)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAnguilla.RegisterName(xlanguage.Spanish, "Anguila")
	dataAnguilla.RegisterOfficialName(xlanguage.Spanish, "Anguila")
	dataAnguilla.RegisterCapital(xlanguage.Spanish, "El Valle")
}
