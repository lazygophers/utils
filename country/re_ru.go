//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataReunion.RegisterName(xlanguage.Russian, "Реюньон")
	dataReunion.RegisterOfficialName(xlanguage.Russian, "Реюньон")
	dataReunion.RegisterCapital(xlanguage.Russian, "Сен-Дени")
}
