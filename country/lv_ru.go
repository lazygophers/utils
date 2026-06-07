//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataLatvia.RegisterName(xlanguage.Russian, "Латвия")
	dataLatvia.RegisterOfficialName(xlanguage.Russian, "Латвийская Республика")
	dataLatvia.RegisterCapital(xlanguage.Russian, "Рига")
}
