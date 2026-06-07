//go:build (lang_ru || lang_all) && (country_all || country_americas || country_northern_america || country_pm)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSaintPierreAndMiquelon.RegisterName(xlanguage.Russian, "Сен-Пьер и Микелон")
	dataSaintPierreAndMiquelon.RegisterOfficialName(xlanguage.Russian, "Заморская община Сен-Пьер и Микелон")
	dataSaintPierreAndMiquelon.RegisterCapital(xlanguage.Russian, "Сен-Пьер")
}
