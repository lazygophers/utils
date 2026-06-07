//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCuracao.RegisterName(xlanguage.Russian, "Кюрасао")
	dataCuracao.RegisterOfficialName(xlanguage.Russian, "Страна Кюрасао")
	dataCuracao.RegisterCapital(xlanguage.Russian, "Виллемстад")
}
