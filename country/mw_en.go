//go:build country_africa || country_all || country_eastern_africa || country_mw

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMalawi.RegisterName(xlanguage.English, "Malawi")
	dataMalawi.RegisterOfficialName(xlanguage.English, "Republic of Malawi")
	dataMalawi.RegisterCapital(xlanguage.English, "Lilongwe")
}
