//go:build (lang_ru || lang_all) && (country_all || country_americas || country_central_america || country_mx)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMexico.RegisterName(xlanguage.Russian, "Мексика")
	dataMexico.RegisterOfficialName(xlanguage.Russian, "Мексиканские Соединённые Штаты")
	dataMexico.RegisterCapital(xlanguage.Russian, "Мехико")
}
