//go:build country_all || country_americas || country_central_america || country_hn

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataHonduras.RegisterName(xlanguage.English, "Honduras")
	dataHonduras.RegisterOfficialName(xlanguage.English, "Republic of Honduras")
	dataHonduras.RegisterCapital(xlanguage.English, "Tegucigalpa")
}
