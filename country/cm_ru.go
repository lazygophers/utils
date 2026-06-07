//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCameroon.RegisterName(xlanguage.Russian, "Камерун")
	dataCameroon.RegisterOfficialName(xlanguage.Russian, "Республика Камерун")
	dataCameroon.RegisterCapital(xlanguage.Russian, "Яунде")
}
