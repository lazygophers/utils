//go:build (lang_ru || lang_all) && (country_all || country_americas || country_caribbean || country_kn)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSaintKittsAndNevis.RegisterName(xlanguage.Russian, "Сент-Китс и Невис")
	dataSaintKittsAndNevis.RegisterOfficialName(xlanguage.Russian, "Федерация Сент-Кристофер и Невис")
	dataSaintKittsAndNevis.RegisterCapital(xlanguage.Russian, "Бастер")
}
