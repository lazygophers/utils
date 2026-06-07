//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSouthKorea.RegisterName(xlanguage.Russian, "Республика Корея")
	dataSouthKorea.RegisterOfficialName(xlanguage.Russian, "Республика Корея")
	dataSouthKorea.RegisterCapital(xlanguage.Russian, "Сеул")
}
