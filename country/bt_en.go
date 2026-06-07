//go:build country_all || country_asia || country_bt || country_southern_asia

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBhutan.RegisterName(xlanguage.English, "Bhutan")
	dataBhutan.RegisterOfficialName(xlanguage.English, "Kingdom of Bhutan")
	dataBhutan.RegisterCapital(xlanguage.English, "Thimphu")
}
