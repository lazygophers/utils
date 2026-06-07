//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSintMaarten.RegisterName(xlanguage.Russian, "Синт-Мартен")
	dataSintMaarten.RegisterOfficialName(xlanguage.Russian, "Синт-Мартен")
	dataSintMaarten.RegisterCapital(xlanguage.Russian, "Филипсбург")
}
