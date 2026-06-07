//go:build (lang_ru || lang_all) && (country_africa || country_all || country_middle_africa || country_st)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSaoTomeAndPrincipe.RegisterName(xlanguage.Russian, "Сан-Томе и Принсипи")
	dataSaoTomeAndPrincipe.RegisterOfficialName(xlanguage.Russian, "Демократическая Республика Сан-Томе и Принсипи")
	dataSaoTomeAndPrincipe.RegisterCapital(xlanguage.Russian, "Сан-Томе")
}
