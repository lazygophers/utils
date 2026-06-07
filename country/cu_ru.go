//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCuba.RegisterName(xlanguage.Russian, "Куба")
	dataCuba.RegisterOfficialName(xlanguage.Russian, "Республика Куба")
	dataCuba.RegisterCapital(xlanguage.Russian, "Гавана")
}
