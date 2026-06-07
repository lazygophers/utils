//go:build (lang_ru || lang_all) && (country_africa || country_all || country_cm || country_middle_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCameroon.RegisterName(xlanguage.Russian, "Камерун")
	dataCameroon.RegisterOfficialName(xlanguage.Russian, "Республика Камерун")
	dataCameroon.RegisterCapital(xlanguage.Russian, "Яунде")
}
