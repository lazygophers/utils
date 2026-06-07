package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataVenezuela.RegisterName(xlanguage.Spanish, "Venezuela")
	dataVenezuela.RegisterOfficialName(xlanguage.Spanish, "República Bolivariana de Venezuela")
	dataVenezuela.RegisterCapital(xlanguage.Spanish, "Caracas")
}
