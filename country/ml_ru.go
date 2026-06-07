//go:build (lang_ru || lang_all) && (country_africa || country_all || country_ml || country_western_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMali.RegisterName(xlanguage.Russian, "Мали")
	dataMali.RegisterOfficialName(xlanguage.Russian, "Республика Мали")
	dataMali.RegisterCapital(xlanguage.Russian, "Бамако")
}
