//go:build country_africa || country_all || country_eastern_africa || country_zw

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataZimbabwe.RegisterName(xlanguage.English, "Zimbabwe")
	dataZimbabwe.RegisterOfficialName(xlanguage.English, "Republic of Zimbabwe")
	dataZimbabwe.RegisterCapital(xlanguage.English, "Harare")
}
