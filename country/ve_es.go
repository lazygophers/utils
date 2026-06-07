//go:build country_all || country_americas || country_south_america || country_ve

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataVenezuela.RegisterName(xlanguage.Spanish, "Venezuela")
	dataVenezuela.RegisterOfficialName(xlanguage.Spanish, "República Bolivariana de Venezuela")
	dataVenezuela.RegisterCapital(xlanguage.Spanish, "Caracas")
}
