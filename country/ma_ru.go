//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMorocco.RegisterName(xlanguage.Russian, "Марокко")
	dataMorocco.RegisterOfficialName(xlanguage.Russian, "Королевство Марокко")
	dataMorocco.RegisterCapital(xlanguage.Russian, "Рабат")
}
