//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataVanuatu.RegisterName(xlanguage.Russian, "Вануату")
	dataVanuatu.RegisterOfficialName(xlanguage.Russian, "Республика Вануату")
	dataVanuatu.RegisterCapital(xlanguage.Russian, "Порт-Вила")
}
