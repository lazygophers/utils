//go:build (lang_ru || lang_all) && (country_all || country_asia || country_pk || country_southern_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataPakistan.RegisterName(xlanguage.Russian, "Пакистан")
	dataPakistan.RegisterOfficialName(xlanguage.Russian, "Исламская Республика Пакистан")
	dataPakistan.RegisterCapital(xlanguage.Russian, "Исламабад")
}
