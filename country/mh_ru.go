//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMarshallIslands.RegisterName(xlanguage.Russian, "Маршалловы Острова")
	dataMarshallIslands.RegisterOfficialName(xlanguage.Russian, "Республика Маршалловы Острова")
	dataMarshallIslands.RegisterCapital(xlanguage.Russian, "Маджуро")
}
