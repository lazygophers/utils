//go:build country_all || country_americas || country_central_america || country_hn

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataHonduras.RegisterName(xlanguage.Spanish, "Honduras")
	dataHonduras.RegisterOfficialName(xlanguage.Spanish, "República de Honduras")
	dataHonduras.RegisterCapital(xlanguage.Spanish, "Tegucigalpa")
}
