//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAlgeria.RegisterName(xlanguage.Russian, "Алжир")
	dataAlgeria.RegisterOfficialName(xlanguage.Russian, "Алжирская Народная Демократическая Республика")
	dataAlgeria.RegisterCapital(xlanguage.Russian, "Алжир")
}
