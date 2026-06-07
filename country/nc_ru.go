//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNewCaledonia.RegisterName(xlanguage.Russian, "Новая Каледония")
	dataNewCaledonia.RegisterOfficialName(xlanguage.Russian, "Новая Каледония")
	dataNewCaledonia.RegisterCapital(xlanguage.Russian, "Нумеа")
}
