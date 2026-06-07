//go:build (lang_es || lang_all) && (country_africa || country_all || country_bf || country_western_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBurkinaFaso.RegisterName(xlanguage.Spanish, "Burkina Faso")
	dataBurkinaFaso.RegisterOfficialName(xlanguage.Spanish, "Burkina Faso")
	dataBurkinaFaso.RegisterCapital(xlanguage.Spanish, "Uagadugú")
}
