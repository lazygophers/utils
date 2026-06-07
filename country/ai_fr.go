//go:build (lang_fr || lang_all) && (country_ai || country_all || country_americas || country_caribbean)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAnguilla.RegisterName(xlanguage.French, "Anguilla")
	dataAnguilla.RegisterOfficialName(xlanguage.French, "Anguilla")
	dataAnguilla.RegisterCapital(xlanguage.French, "The Valley")
}
