//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSaintBarthelemy.RegisterName(xlanguage.Russian, "Сен-Бартелеми")
	dataSaintBarthelemy.RegisterOfficialName(xlanguage.Russian, "Заморская община Сен-Бартелеми")
	dataSaintBarthelemy.RegisterCapital(xlanguage.Russian, "Густавия")
}
