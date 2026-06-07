//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataHonduras.RegisterName(xlanguage.Russian, "Гондурас")
	dataHonduras.RegisterOfficialName(xlanguage.Russian, "Республика Гондурас")
	dataHonduras.RegisterCapital(xlanguage.Russian, "Тегусигальпа")
}
