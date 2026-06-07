//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataUganda.RegisterName(xlanguage.Russian, "Уганда")
	dataUganda.RegisterOfficialName(xlanguage.Russian, "Республика Уганда")
	dataUganda.RegisterCapital(xlanguage.Russian, "Кампала")
}
