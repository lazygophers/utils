//go:build country_all || country_americas || country_ec || country_south_america

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataEcuador.RegisterName(xlanguage.English, "Ecuador")
	dataEcuador.RegisterOfficialName(xlanguage.English, "Republic of Ecuador")
	dataEcuador.RegisterCapital(xlanguage.English, "Quito")
}
