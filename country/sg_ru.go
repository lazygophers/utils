//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSingapore.RegisterName(xlanguage.Russian, "Сингапур")
	dataSingapore.RegisterOfficialName(xlanguage.Russian, "Республика Сингапур")
	dataSingapore.RegisterCapital(xlanguage.Russian, "Сингапур")
}
