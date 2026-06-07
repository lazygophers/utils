//go:build country_africa || country_all || country_cm || country_middle_africa

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCameroon.RegisterName(xlanguage.English, "Cameroon")
	dataCameroon.RegisterOfficialName(xlanguage.English, "Republic of Cameroon")
	dataCameroon.RegisterCapital(xlanguage.English, "Yaounde")
}
