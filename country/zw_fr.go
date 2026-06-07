//go:build (lang_fr || lang_all) && (country_africa || country_all || country_eastern_africa || country_zw)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataZimbabwe.RegisterName(xlanguage.French, "Zimbabwe")
	dataZimbabwe.RegisterOfficialName(xlanguage.French, "République du Zimbabwe")
	dataZimbabwe.RegisterCapital(xlanguage.French, "Harare")
}
