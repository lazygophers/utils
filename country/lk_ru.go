//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSriLanka.RegisterName(xlanguage.Russian, "Шри-Ланка")
	dataSriLanka.RegisterOfficialName(xlanguage.Russian, "Демократическая Социалистическая Республика Шри-Ланка")
	dataSriLanka.RegisterCapital(xlanguage.Russian, "Шри-Джаяварденепура-Котте")
}
