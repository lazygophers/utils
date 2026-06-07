//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMonaco.RegisterName(xlanguage.Spanish, "Mónaco")
	dataMonaco.RegisterOfficialName(xlanguage.Spanish, "Principado de Mónaco")
	dataMonaco.RegisterCapital(xlanguage.Spanish, "Mónaco")
}
