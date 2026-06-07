//go:build (lang_ru || lang_all) && (country_all || country_europe || country_gg || country_northern_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGuernsey.RegisterName(xlanguage.Russian, "Гернси")
	dataGuernsey.RegisterOfficialName(xlanguage.Russian, "Бейливик Гернси")
	dataGuernsey.RegisterCapital(xlanguage.Russian, "Сент-Питер-Порт")
}
