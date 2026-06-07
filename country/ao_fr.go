//go:build (lang_fr || lang_all) && (country_africa || country_all || country_ao || country_middle_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAngola.RegisterName(xlanguage.French, "Angola")
	dataAngola.RegisterOfficialName(xlanguage.French, "République d'Angola")
	dataAngola.RegisterCapital(xlanguage.French, "Luanda")
}
