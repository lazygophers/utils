//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBangladesh.RegisterName(xlanguage.Russian, "Бангладеш")
	dataBangladesh.RegisterOfficialName(xlanguage.Russian, "Народная Республика Бангладеш")
	dataBangladesh.RegisterCapital(xlanguage.Russian, "Дакка")
}
