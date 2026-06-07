//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataChile.RegisterName(xlanguage.Russian, "Чили")
	dataChile.RegisterOfficialName(xlanguage.Russian, "Республика Чили")
	dataChile.RegisterCapital(xlanguage.Russian, "Сантьяго")
}
