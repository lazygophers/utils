//go:build (lang_ru || lang_all) && (country_all || country_americas || country_caribbean || country_mf)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSaintMartin.RegisterName(xlanguage.Russian, "Сен-Мартен")
	dataSaintMartin.RegisterOfficialName(xlanguage.Russian, "Заморская община Сен-Мартен")
	dataSaintMartin.RegisterCapital(xlanguage.Russian, "Мариго")
}
