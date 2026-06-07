//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataPitcairn.RegisterName(xlanguage.Russian, "Острова Питкэрн")
	dataPitcairn.RegisterOfficialName(xlanguage.Russian, "Питкэрн, Хендерсон, Дюси и Оэно")
	dataPitcairn.RegisterCapital(xlanguage.Russian, "Адамстаун")
}
