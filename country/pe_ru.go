//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataPeru.RegisterName(xlanguage.Russian, "Перу")
	dataPeru.RegisterOfficialName(xlanguage.Russian, "Республика Перу")
	dataPeru.RegisterCapital(xlanguage.Russian, "Лима")
}
