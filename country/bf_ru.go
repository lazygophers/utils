//go:build (lang_ru || lang_all) && (country_africa || country_all || country_bf || country_western_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBurkinaFaso.RegisterName(xlanguage.Russian, "Буркина-Фасо")
	dataBurkinaFaso.RegisterOfficialName(xlanguage.Russian, "Буркина-Фасо")
	dataBurkinaFaso.RegisterCapital(xlanguage.Russian, "Уагадугу")
}
