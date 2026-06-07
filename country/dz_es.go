//go:build (lang_es || lang_all) && (country_africa || country_all || country_dz || country_northern_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAlgeria.RegisterName(xlanguage.Spanish, "Argelia")
	dataAlgeria.RegisterOfficialName(xlanguage.Spanish, "República Argelina Democrática y Popular")
	dataAlgeria.RegisterCapital(xlanguage.Spanish, "Argel")
}
