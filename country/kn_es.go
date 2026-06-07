//go:build (lang_es || lang_all) && (country_all || country_americas || country_caribbean || country_kn)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSaintKittsAndNevis.RegisterName(xlanguage.Spanish, "San Cristóbal y Nieves")
	dataSaintKittsAndNevis.RegisterOfficialName(xlanguage.Spanish, "Federación de San Cristóbal y Nieves")
	dataSaintKittsAndNevis.RegisterCapital(xlanguage.Spanish, "Basseterre")
}
