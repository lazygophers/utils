//go:build (lang_ru || lang_all) && (country_all || country_fm || country_micronesia || country_oceania)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMicronesia.RegisterName(xlanguage.Russian, "Микронезия")
	dataMicronesia.RegisterOfficialName(xlanguage.Russian, "Федеративные Штаты Микронезии")
	dataMicronesia.RegisterCapital(xlanguage.Russian, "Паликир")
}
