//go:build (lang_ru || lang_all) && (country_all || country_americas || country_cl || country_south_america)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataChile.RegisterName(xlanguage.Russian, "Чили")
	dataChile.RegisterOfficialName(xlanguage.Russian, "Республика Чили")
	dataChile.RegisterCapital(xlanguage.Russian, "Сантьяго")
}
