//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSomalia.RegisterName(xlanguage.Spanish, "Somalia")
	dataSomalia.RegisterOfficialName(xlanguage.Spanish, "República Federal de Somalia")
	dataSomalia.RegisterCapital(xlanguage.Spanish, "Mogadiscio")
}
