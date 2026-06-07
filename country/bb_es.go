//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBarbados.RegisterName(xlanguage.Spanish, "Barbados")
	dataBarbados.RegisterOfficialName(xlanguage.Spanish, "Barbados")
	dataBarbados.RegisterCapital(xlanguage.Spanish, "Bridgetown")
}
