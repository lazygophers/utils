//go:build (lang_ru || lang_all) && (country_all || country_americas || country_bm || country_northern_america)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBermuda.RegisterName(xlanguage.Russian, "Бермудские Острова")
	dataBermuda.RegisterOfficialName(xlanguage.Russian, "Бермудские Острова")
	dataBermuda.RegisterCapital(xlanguage.Russian, "Гамильтон")
}
