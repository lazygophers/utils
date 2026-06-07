//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMali.RegisterName(xlanguage.Russian, "Мали")
	dataMali.RegisterOfficialName(xlanguage.Russian, "Республика Мали")
	dataMali.RegisterCapital(xlanguage.Russian, "Бамако")
}
