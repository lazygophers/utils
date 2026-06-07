//go:build (lang_es || lang_all) && (country_africa || country_all || country_middle_africa || country_st)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSaoTomeAndPrincipe.RegisterName(xlanguage.Spanish, "Santo Tomé y Príncipe")
	dataSaoTomeAndPrincipe.RegisterOfficialName(xlanguage.Spanish, "República Democrática de Santo Tomé y Príncipe")
	dataSaoTomeAndPrincipe.RegisterCapital(xlanguage.Spanish, "Santo Tomé")
}
