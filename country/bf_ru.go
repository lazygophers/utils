//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBurkinaFaso.RegisterName(xlanguage.Russian, "Буркина-Фасо")
	dataBurkinaFaso.RegisterOfficialName(xlanguage.Russian, "Буркина-Фасо")
	dataBurkinaFaso.RegisterCapital(xlanguage.Russian, "Уагадугу")
}
