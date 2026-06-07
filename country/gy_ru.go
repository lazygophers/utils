//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGuyana.RegisterName(xlanguage.Russian, "Гайана")
	dataGuyana.RegisterOfficialName(xlanguage.Russian, "Кооперативная Республика Гайана")
	dataGuyana.RegisterCapital(xlanguage.Russian, "Джорджтаун")
}
