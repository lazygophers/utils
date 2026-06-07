//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMicronesia.RegisterName(xlanguage.Spanish, "Micronesia")
	dataMicronesia.RegisterOfficialName(xlanguage.Spanish, "Estados Federados de Micronesia")
	dataMicronesia.RegisterCapital(xlanguage.Spanish, "Palikir")
}
