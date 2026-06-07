//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataEthiopia.RegisterName(xlanguage.Russian, "Эфиопия")
	dataEthiopia.RegisterOfficialName(xlanguage.Russian, "Федеративная Демократическая Республика Эфиопия")
	dataEthiopia.RegisterCapital(xlanguage.Russian, "Аддис-Абеба")
}
