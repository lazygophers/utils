//go:build (lang_ru || lang_all) && (country_all || country_europe || country_je || country_northern_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataJersey.RegisterName(xlanguage.Russian, "Джерси")
	dataJersey.RegisterOfficialName(xlanguage.Russian, "Бейливик Джерси")
	dataJersey.RegisterCapital(xlanguage.Russian, "Сент-Хелиер")
}
