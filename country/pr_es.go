//go:build country_all || country_americas || country_caribbean || country_pr

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataPuertoRico.RegisterName(xlanguage.Spanish, "Puerto Rico")
	dataPuertoRico.RegisterOfficialName(xlanguage.Spanish, "Estado Libre Asociado de Puerto Rico")
	dataPuertoRico.RegisterCapital(xlanguage.Spanish, "San Juan")
}
