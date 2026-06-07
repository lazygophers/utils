//go:build country_all || country_americas || country_caribbean || country_tt

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTrinidadAndTobago.RegisterName(xlanguage.English, "Trinidad and Tobago")
	dataTrinidadAndTobago.RegisterOfficialName(xlanguage.English, "Republic of Trinidad and Tobago")
	dataTrinidadAndTobago.RegisterCapital(xlanguage.English, "Port of Spain")
}
