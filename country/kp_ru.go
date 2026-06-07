//go:build (lang_ru || lang_all) && (country_all || country_asia || country_eastern_asia || country_kp)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNorthKorea.RegisterName(xlanguage.Russian, "КНДР")
	dataNorthKorea.RegisterOfficialName(xlanguage.Russian, "Корейская Народно-Демократическая Республика")
	dataNorthKorea.RegisterCapital(xlanguage.Russian, "Пхеньян")
}
