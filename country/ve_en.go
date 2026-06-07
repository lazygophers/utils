//go:build country_all || country_americas || country_south_america || country_ve

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataVenezuela.RegisterName(xlanguage.English, "Venezuela")
	dataVenezuela.RegisterOfficialName(xlanguage.English, "Bolivarian Republic of Venezuela")
	dataVenezuela.RegisterCapital(xlanguage.English, "Caracas")
}
