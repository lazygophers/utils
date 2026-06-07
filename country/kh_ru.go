//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCambodia.RegisterName(xlanguage.Russian, "Камбоджа")
	dataCambodia.RegisterOfficialName(xlanguage.Russian, "Королевство Камбоджа")
	dataCambodia.RegisterCapital(xlanguage.Russian, "Пномпень")
}
