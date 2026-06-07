//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataComoros.RegisterName(xlanguage.Russian, "Коморы")
	dataComoros.RegisterOfficialName(xlanguage.Russian, "Союз Коморских Островов")
	dataComoros.RegisterCapital(xlanguage.Russian, "Морони")
}
