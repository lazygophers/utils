//go:build (lang_ru || lang_all) && (country_all || country_europe || country_nl || country_western_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNetherlands.RegisterName(xlanguage.Russian, "Нидерланды")
	dataNetherlands.RegisterOfficialName(xlanguage.Russian, "Королевство Нидерландов")
	dataNetherlands.RegisterCapital(xlanguage.Russian, "Амстердам")
}
