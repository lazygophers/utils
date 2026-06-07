//go:build (lang_es || lang_all) && (country_africa || country_all || country_eastern_africa || country_yt)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMayotte.RegisterName(xlanguage.Spanish, "Mayotte")
	dataMayotte.RegisterOfficialName(xlanguage.Spanish, "Departamento de Mayotte")
	dataMayotte.RegisterCapital(xlanguage.Spanish, "Mamoudzou")
}
