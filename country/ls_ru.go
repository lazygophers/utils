//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataLesotho.RegisterName(xlanguage.Russian, "Лесото")
	dataLesotho.RegisterOfficialName(xlanguage.Russian, "Королевство Лесото")
	dataLesotho.RegisterCapital(xlanguage.Russian, "Масеру")
}
