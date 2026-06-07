//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataVenezuela.RegisterName(xlanguage.Russian, "Венесуэла")
	dataVenezuela.RegisterOfficialName(xlanguage.Russian, "Боливарианская Республика Венесуэла")
	dataVenezuela.RegisterCapital(xlanguage.Russian, "Каракас")
}
