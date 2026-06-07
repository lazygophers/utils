//go:build (lang_es || lang_all) && (country_all || country_americas || country_gf || country_south_america)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataFrenchGuiana.RegisterName(xlanguage.Spanish, "Guayana Francesa")
	dataFrenchGuiana.RegisterOfficialName(xlanguage.Spanish, "Guayana Francesa")
	dataFrenchGuiana.RegisterCapital(xlanguage.Spanish, "Cayena")
}
