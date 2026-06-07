//go:build (lang_ru || lang_all) && (country_ag || country_all || country_americas || country_caribbean)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAntiguaAndBarbuda.RegisterName(xlanguage.Russian, "Антигуа и Барбуда")
	dataAntiguaAndBarbuda.RegisterOfficialName(xlanguage.Russian, "Антигуа и Барбуда")
	dataAntiguaAndBarbuda.RegisterCapital(xlanguage.Russian, "Сент-Джонс")
}
