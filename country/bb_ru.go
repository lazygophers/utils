//go:build (lang_ru || lang_all) && (country_all || country_americas || country_bb || country_caribbean)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBarbados.RegisterName(xlanguage.Russian, "Барбадос")
	dataBarbados.RegisterOfficialName(xlanguage.Russian, "Барбадос")
	dataBarbados.RegisterCapital(xlanguage.Russian, "Бриджтаун")
}
