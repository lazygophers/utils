//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSlovenia.RegisterName(xlanguage.Russian, "Словения")
	dataSlovenia.RegisterOfficialName(xlanguage.Russian, "Республика Словения")
	dataSlovenia.RegisterCapital(xlanguage.Russian, "Любляна")
}
