//go:build (lang_es || lang_all) && (country_africa || country_all || country_cm || country_middle_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCameroon.RegisterName(xlanguage.Spanish, "Camerún")
	dataCameroon.RegisterOfficialName(xlanguage.Spanish, "República de Camerún")
	dataCameroon.RegisterCapital(xlanguage.Spanish, "Yaundé")
}
