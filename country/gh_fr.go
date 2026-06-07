//go:build (lang_fr || lang_all) && (country_africa || country_all || country_gh || country_western_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGhana.RegisterName(xlanguage.French, "Ghana")
	dataGhana.RegisterOfficialName(xlanguage.French, "République du Ghana")
	dataGhana.RegisterCapital(xlanguage.French, "Accra")
}
