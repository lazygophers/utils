//go:build (lang_ru || lang_all) && (country_africa || country_all || country_eastern_africa || country_zw)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataZimbabwe.RegisterName(xlanguage.Russian, "Зимбабве")
	dataZimbabwe.RegisterOfficialName(xlanguage.Russian, "Республика Зимбабве")
	dataZimbabwe.RegisterCapital(xlanguage.Russian, "Хараре")
}
