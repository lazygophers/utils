//go:build country_all || country_americas || country_caribbean || country_ms

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMontserrat.RegisterName(xlanguage.English, "Montserrat")
	dataMontserrat.RegisterOfficialName(xlanguage.English, "Montserrat")
	dataMontserrat.RegisterCapital(xlanguage.English, "Plymouth")
}
