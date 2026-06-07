//go:build (lang_es || lang_all) && (country_africa || country_all || country_middle_africa || country_td)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataChad.RegisterName(xlanguage.Spanish, "Chad")
	dataChad.RegisterOfficialName(xlanguage.Spanish, "República del Chad")
	dataChad.RegisterCapital(xlanguage.Spanish, "Yamena")
}
