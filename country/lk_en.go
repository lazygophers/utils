//go:build country_all || country_asia || country_lk || country_southern_asia

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSriLanka.RegisterName(xlanguage.English, "Sri Lanka")
	dataSriLanka.RegisterOfficialName(xlanguage.English, "Democratic Socialist Republic of Sri Lanka")
	dataSriLanka.RegisterCapital(xlanguage.English, "Sri Jayawardenepura Kotte")
}
