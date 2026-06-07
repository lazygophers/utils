//go:build (lang_es || lang_all) && (country_all || country_asia || country_kh || country_south_eastern_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCambodia.RegisterName(xlanguage.Spanish, "Camboya")
	dataCambodia.RegisterOfficialName(xlanguage.Spanish, "Reino de Camboya")
	dataCambodia.RegisterCapital(xlanguage.Spanish, "Nom Pen")
}
