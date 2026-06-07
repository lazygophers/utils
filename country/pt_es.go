//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataPortugal.RegisterName(xlanguage.Spanish, "Portugal")
	dataPortugal.RegisterOfficialName(xlanguage.Spanish, "República Portuguesa")
	dataPortugal.RegisterCapital(xlanguage.Spanish, "Lisboa")
}
