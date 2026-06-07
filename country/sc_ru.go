//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSeychelles.RegisterName(xlanguage.Russian, "Сейшельские Острова")
	dataSeychelles.RegisterOfficialName(xlanguage.Russian, "Республика Сейшельские Острова")
	dataSeychelles.RegisterCapital(xlanguage.Russian, "Виктория")
}
