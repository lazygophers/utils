//go:build (lang_ru || lang_all) && (country_africa || country_all || country_cf || country_middle_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCentralAfricanRepublic.RegisterName(xlanguage.Russian, "Центральноафриканская Республика")
	dataCentralAfricanRepublic.RegisterOfficialName(xlanguage.Russian, "Центральноафриканская Республика")
	dataCentralAfricanRepublic.RegisterCapital(xlanguage.Russian, "Банги")
}
