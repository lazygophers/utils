//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMontserrat.RegisterName(xlanguage.Russian, "Монтсеррат")
	dataMontserrat.RegisterOfficialName(xlanguage.Russian, "Монтсеррат")
	dataMontserrat.RegisterCapital(xlanguage.Russian, "Плимут")
}
