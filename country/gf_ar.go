//go:build (lang_ar || lang_all) && (country_all || country_americas || country_gf || country_south_america)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataFrenchGuiana.RegisterName(xlanguage.Arabic, "غويانا الفرنسية")
	dataFrenchGuiana.RegisterOfficialName(xlanguage.Arabic, "غويانا الفرنسية")
	dataFrenchGuiana.RegisterCapital(xlanguage.Arabic, "كايين")
}
