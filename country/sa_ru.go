//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSaudiArabia.RegisterName(xlanguage.Russian, "Саудовская Аравия")
	dataSaudiArabia.RegisterOfficialName(xlanguage.Russian, "Королевство Саудовская Аравия")
	dataSaudiArabia.RegisterCapital(xlanguage.Russian, "Эр-Рияд")
}
