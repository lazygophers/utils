//go:build (lang_es || lang_all) && (country_africa || country_all || country_ao || country_middle_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAngola.RegisterName(xlanguage.Spanish, "Angola")
	dataAngola.RegisterOfficialName(xlanguage.Spanish, "República de Angola")
	dataAngola.RegisterCapital(xlanguage.Spanish, "Luanda")
}
