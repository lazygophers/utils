//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBhutan.RegisterName(xlanguage.Russian, "Бутан")
	dataBhutan.RegisterOfficialName(xlanguage.Russian, "Королевство Бутан")
	dataBhutan.RegisterCapital(xlanguage.Russian, "Тхимпху")
}
