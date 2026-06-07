//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataThailand.RegisterName(xlanguage.Russian, "Таиланд")
	dataThailand.RegisterOfficialName(xlanguage.Russian, "Королевство Таиланд")
	dataThailand.RegisterCapital(xlanguage.Russian, "Бангкок")
}
