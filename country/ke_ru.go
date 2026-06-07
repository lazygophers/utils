//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataKenya.RegisterName(xlanguage.Russian, "Кения")
	dataKenya.RegisterOfficialName(xlanguage.Russian, "Республика Кения")
	dataKenya.RegisterCapital(xlanguage.Russian, "Найроби")
}
