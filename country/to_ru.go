//go:build (lang_ru || lang_all) && (country_all || country_oceania || country_polynesia || country_to)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTonga.RegisterName(xlanguage.Russian, "Тонга")
	dataTonga.RegisterOfficialName(xlanguage.Russian, "Королевство Тонга")
	dataTonga.RegisterCapital(xlanguage.Russian, "Нукуалофа")
}
