//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataIvoryCoast.RegisterName(xlanguage.Russian, "Кот-д’Ивуар")
	dataIvoryCoast.RegisterOfficialName(xlanguage.Russian, "Республика Кот-д’Ивуар")
	dataIvoryCoast.RegisterCapital(xlanguage.Russian, "Ямусукро")
}
