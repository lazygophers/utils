//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataFaroeIslands.RegisterName(xlanguage.Spanish, "Islas Feroe")
	dataFaroeIslands.RegisterOfficialName(xlanguage.Spanish, "Islas Feroe")
	dataFaroeIslands.RegisterCapital(xlanguage.Spanish, "Tórshavn")
}
