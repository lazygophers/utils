//go:build (lang_fr || lang_all) && (country_all || country_americas || country_south_america || country_ve)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataVenezuela.RegisterName(xlanguage.French, "Venezuela")
	dataVenezuela.RegisterOfficialName(xlanguage.French, "République bolivarienne du Venezuela")
	dataVenezuela.RegisterCapital(xlanguage.French, "Caracas")
}
