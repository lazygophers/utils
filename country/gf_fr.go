//go:build country_all || country_americas || country_gf || country_south_america

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataFrenchGuiana.RegisterName(xlanguage.French, "Guyane")
	dataFrenchGuiana.RegisterOfficialName(xlanguage.French, "Guyane française")
	dataFrenchGuiana.RegisterCapital(xlanguage.French, "Cayenne")
}
