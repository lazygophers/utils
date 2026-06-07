//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNetherlands.RegisterName(xlanguage.Russian, "Нидерланды")
	dataNetherlands.RegisterOfficialName(xlanguage.Russian, "Королевство Нидерландов")
	dataNetherlands.RegisterCapital(xlanguage.Russian, "Амстердам")
}
