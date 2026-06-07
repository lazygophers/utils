//go:build (lang_ru || lang_all) && (country_africa || country_all || country_ma || country_northern_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMorocco.RegisterName(xlanguage.Russian, "Марокко")
	dataMorocco.RegisterOfficialName(xlanguage.Russian, "Королевство Марокко")
	dataMorocco.RegisterCapital(xlanguage.Russian, "Рабат")
}
