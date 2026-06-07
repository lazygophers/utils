//go:build (lang_es || lang_all) && (country_africa || country_all || country_eastern_africa || country_zw)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataZimbabwe.RegisterName(xlanguage.Spanish, "Zimbabue")
	dataZimbabwe.RegisterOfficialName(xlanguage.Spanish, "República de Zimbabue")
	dataZimbabwe.RegisterCapital(xlanguage.Spanish, "Harare")
}
