//go:build (lang_ja || lang_all) && (country_all || country_americas || country_gf || country_south_america)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataFrenchGuiana.RegisterName(xlanguage.Japanese, "フランス領ギアナ")
	dataFrenchGuiana.RegisterOfficialName(xlanguage.Japanese, "フランス領ギアナ")
	dataFrenchGuiana.RegisterCapital(xlanguage.Japanese, "カイエンヌ")
}
