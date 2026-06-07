//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBritishIndianOceanTerritory.RegisterName(xlanguage.Russian, "Британская территория в Индийском океане")
	dataBritishIndianOceanTerritory.RegisterOfficialName(xlanguage.Russian, "Британская территория в Индийском океане")
	dataBritishIndianOceanTerritory.RegisterCapital(xlanguage.Russian, "Диего-Гарсия")
}
