//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSwitzerland.RegisterName(xlanguage.Spanish, "Suiza")
	dataSwitzerland.RegisterOfficialName(xlanguage.Spanish, "Confederación Suiza")
	dataSwitzerland.RegisterCapital(xlanguage.Spanish, "Berna")
}
