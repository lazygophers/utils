//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCuba.RegisterName(xlanguage.French, "Cuba")
	dataCuba.RegisterOfficialName(xlanguage.French, "République de Cuba")
	dataCuba.RegisterCapital(xlanguage.French, "La Havane")
}
