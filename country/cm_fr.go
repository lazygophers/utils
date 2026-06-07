//go:build country_africa || country_all || country_cm || country_middle_africa

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCameroon.RegisterName(xlanguage.French, "Cameroun")
	dataCameroon.RegisterOfficialName(xlanguage.French, "République du Cameroun")
	dataCameroon.RegisterCapital(xlanguage.French, "Yaoundé")
}
