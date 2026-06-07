//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCroatia.RegisterName(xlanguage.Russian, "Хорватия")
	dataCroatia.RegisterOfficialName(xlanguage.Russian, "Республика Хорватия")
	dataCroatia.RegisterCapital(xlanguage.Russian, "Загреб")
}
