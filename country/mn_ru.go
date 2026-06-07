//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMongolia.RegisterName(xlanguage.Russian, "Монголия")
	dataMongolia.RegisterOfficialName(xlanguage.Russian, "Монголия")
	dataMongolia.RegisterCapital(xlanguage.Russian, "Улан-Батор")
}
