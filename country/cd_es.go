//go:build (lang_es || lang_all) && (country_africa || country_all || country_cd || country_middle_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataDrCongo.RegisterName(xlanguage.Spanish, "República Democrática del Congo")
	dataDrCongo.RegisterOfficialName(xlanguage.Spanish, "República Democrática del Congo")
	dataDrCongo.RegisterCapital(xlanguage.Spanish, "Kinsasa")
}
