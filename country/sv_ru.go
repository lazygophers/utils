//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataElSalvador.RegisterName(xlanguage.Russian, "Сальвадор")
	dataElSalvador.RegisterOfficialName(xlanguage.Russian, "Республика Эль-Сальвадор")
	dataElSalvador.RegisterCapital(xlanguage.Russian, "Сан-Сальвадор")
}
