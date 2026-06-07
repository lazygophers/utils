//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataEquatorialGuinea.RegisterName(xlanguage.Russian, "Экваториальная Гвинея")
	dataEquatorialGuinea.RegisterOfficialName(xlanguage.Russian, "Республика Экваториальная Гвинея")
	dataEquatorialGuinea.RegisterCapital(xlanguage.Russian, "Малабо")
}
