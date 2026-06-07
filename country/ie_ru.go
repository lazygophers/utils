//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataIreland.RegisterName(xlanguage.Russian, "Ирландия")
	dataIreland.RegisterOfficialName(xlanguage.Russian, "Республика Ирландия")
	dataIreland.RegisterCapital(xlanguage.Russian, "Дублин")
}
