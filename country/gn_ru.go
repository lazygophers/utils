//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGuinea.RegisterName(xlanguage.Russian, "Гвинея")
	dataGuinea.RegisterOfficialName(xlanguage.Russian, "Гвинейская Республика")
	dataGuinea.RegisterCapital(xlanguage.Russian, "Конакри")
}
