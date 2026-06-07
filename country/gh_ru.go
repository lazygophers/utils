//go:build (lang_ru || lang_all) && (country_africa || country_all || country_gh || country_western_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGhana.RegisterName(xlanguage.Russian, "Гана")
	dataGhana.RegisterOfficialName(xlanguage.Russian, "Республика Гана")
	dataGhana.RegisterCapital(xlanguage.Russian, "Аккра")
}
