//go:build country_all || country_americas || country_co || country_south_america

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataColombia.RegisterName(xlanguage.Spanish, "Colombia")
	dataColombia.RegisterOfficialName(xlanguage.Spanish, "República de Colombia")
	dataColombia.RegisterCapital(xlanguage.Spanish, "Bogotá")
}
