//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataEstonia.RegisterName(xlanguage.Russian, "Эстония")
	dataEstonia.RegisterOfficialName(xlanguage.Russian, "Эстонская Республика")
	dataEstonia.RegisterCapital(xlanguage.Russian, "Таллин")
}
