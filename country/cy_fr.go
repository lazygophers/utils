//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCyprus.RegisterName(xlanguage.French, "Chypre")
	dataCyprus.RegisterOfficialName(xlanguage.French, "République de Chypre")
	dataCyprus.RegisterCapital(xlanguage.French, "Nicosie")
}
