//go:build (lang_fr || lang_all) && (country_all || country_americas || country_caribbean || country_ms)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMontserrat.RegisterName(xlanguage.French, "Montserrat")
	dataMontserrat.RegisterOfficialName(xlanguage.French, "Montserrat")
	dataMontserrat.RegisterCapital(xlanguage.French, "Plymouth")
}
