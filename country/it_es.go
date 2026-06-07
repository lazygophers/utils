//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataItaly.RegisterName(xlanguage.Spanish, "Italia")
	dataItaly.RegisterOfficialName(xlanguage.Spanish, "República Italiana")
	dataItaly.RegisterCapital(xlanguage.Spanish, "Roma")
}
