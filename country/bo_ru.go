//go:build (lang_ru || lang_all) && (country_all || country_americas || country_bo || country_south_america)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBolivia.RegisterName(xlanguage.Russian, "Боливия")
	dataBolivia.RegisterOfficialName(xlanguage.Russian, "Многонациональное Государство Боливия")
	dataBolivia.RegisterCapital(xlanguage.Russian, "Сукре")
}
