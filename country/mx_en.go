//go:build country_all || country_americas || country_central_america || country_mx

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMexico.RegisterName(xlanguage.English, "Mexico")
	dataMexico.RegisterOfficialName(xlanguage.English, "United Mexican States")
	dataMexico.RegisterCapital(xlanguage.English, "Mexico City")
}
