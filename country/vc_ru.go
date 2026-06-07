//go:build (lang_ru || lang_all) && (country_all || country_americas || country_caribbean || country_vc)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSaintVincentAndGrenadines.RegisterName(xlanguage.Russian, "Сент-Винсент и Гренадины")
	dataSaintVincentAndGrenadines.RegisterOfficialName(xlanguage.Russian, "Сент-Винсент и Гренадины")
	dataSaintVincentAndGrenadines.RegisterCapital(xlanguage.Russian, "Кингстаун")
}
