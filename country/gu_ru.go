//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGuam.RegisterName(xlanguage.Russian, "Гуам")
	dataGuam.RegisterOfficialName(xlanguage.Russian, "Территория Гуам")
	dataGuam.RegisterCapital(xlanguage.Russian, "Хагатна")
}
