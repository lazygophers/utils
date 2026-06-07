//go:build country_all || country_americas || country_ec || country_south_america

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataEcuador.RegisterName(xlanguage.Spanish, "Ecuador")
	dataEcuador.RegisterOfficialName(xlanguage.Spanish, "República del Ecuador")
	dataEcuador.RegisterCapital(xlanguage.Spanish, "Quito")
}
