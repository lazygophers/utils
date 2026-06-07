//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCanada.RegisterName(xlanguage.Russian, "Канада")
	dataCanada.RegisterOfficialName(xlanguage.Russian, "Канада")
	dataCanada.RegisterCapital(xlanguage.Russian, "Оттава")
}
