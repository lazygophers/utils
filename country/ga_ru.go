//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGabon.RegisterName(xlanguage.Russian, "Габон")
	dataGabon.RegisterOfficialName(xlanguage.Russian, "Габонская Республика")
	dataGabon.RegisterCapital(xlanguage.Russian, "Либревиль")
}
