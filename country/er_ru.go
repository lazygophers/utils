//go:build (lang_ru || lang_all) && (country_africa || country_all || country_eastern_africa || country_er)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataEritrea.RegisterName(xlanguage.Russian, "Эритрея")
	dataEritrea.RegisterOfficialName(xlanguage.Russian, "Государство Эритрея")
	dataEritrea.RegisterCapital(xlanguage.Russian, "Асмэра")
}
