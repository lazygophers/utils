//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSaintMartin.RegisterName(xlanguage.Russian, "Сен-Мартен")
	dataSaintMartin.RegisterOfficialName(xlanguage.Russian, "Заморская община Сен-Мартен")
	dataSaintMartin.RegisterCapital(xlanguage.Russian, "Мариго")
}
