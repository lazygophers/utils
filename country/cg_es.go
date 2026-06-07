//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCongo.RegisterName(xlanguage.Spanish, "República del Congo")
	dataCongo.RegisterOfficialName(xlanguage.Spanish, "República del Congo")
	dataCongo.RegisterCapital(xlanguage.Spanish, "Brazzaville")
}
