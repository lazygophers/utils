//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMauritius.RegisterName(xlanguage.Russian, "Маврикий")
	dataMauritius.RegisterOfficialName(xlanguage.Russian, "Республика Маврикий")
	dataMauritius.RegisterCapital(xlanguage.Russian, "Порт-Луи")
}
