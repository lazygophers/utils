//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAlandIslands.RegisterName(xlanguage.Russian, "Аландские острова")
	dataAlandIslands.RegisterOfficialName(xlanguage.Russian, "Аландские острова")
	dataAlandIslands.RegisterCapital(xlanguage.Russian, "Мариехамн")
}
