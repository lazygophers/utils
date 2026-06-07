//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataFalklandIslands.RegisterName(xlanguage.Spanish, "Islas Malvinas")
	dataFalklandIslands.RegisterOfficialName(xlanguage.Spanish, "Islas Malvinas")
	dataFalklandIslands.RegisterCapital(xlanguage.Spanish, "Puerto Argentino/Stanley")
}
