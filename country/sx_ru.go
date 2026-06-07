//go:build (lang_ru || lang_all) && (country_all || country_americas || country_caribbean || country_sx)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSintMaarten.RegisterName(xlanguage.Russian, "Синт-Мартен")
	dataSintMaarten.RegisterOfficialName(xlanguage.Russian, "Синт-Мартен")
	dataSintMaarten.RegisterCapital(xlanguage.Russian, "Филипсбург")
}
