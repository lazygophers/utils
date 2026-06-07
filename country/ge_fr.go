//go:build (lang_fr || lang_all) && (country_all || country_asia || country_ge || country_western_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGeorgia.RegisterName(xlanguage.French, "Géorgie")
	dataGeorgia.RegisterOfficialName(xlanguage.French, "Géorgie")
	dataGeorgia.RegisterCapital(xlanguage.French, "Tbilissi")
}
