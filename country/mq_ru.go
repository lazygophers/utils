//go:build (lang_ru || lang_all) && (country_all || country_americas || country_caribbean || country_mq)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMartinique.RegisterName(xlanguage.Russian, "Мартиника")
	dataMartinique.RegisterOfficialName(xlanguage.Russian, "Мартиника")
	dataMartinique.RegisterCapital(xlanguage.Russian, "Фор-де-Франс")
}
